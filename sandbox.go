package main

import (
	"maps"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strings"
)

var aliasRegex = regexp.MustCompile(`alias ([^\ ]+) ([^\ ]+)`)

type Env struct {
	Aliases map[string]string
	Functions map[string]string
	Path []string
}

type Relations struct {
	Aliases []string
	Commands []string
}

func shellCmd(cmd string) string {
	command := exec.Command("fish", "-Nc", cmd)
	command.Env = append(command.Env, "DISABLE_LAZY=1")
	out, err := command.Output()
	handle(err)
	return strings.Trim(string(out), "\n")
}

func loadRelations(cmd string) (relations Relations) {
	before := loadEnv("")
	after := loadEnv(cmd + " && ")
	diff := compareEnvs(before, after)

	aliasSet := make(map[string]bool)
	for alias, _ := range diff.Aliases {
		aliasSet[alias] = true
	}
	for function, _ := range diff.Functions {
		aliasSet[function] = true
	}
	relations.Aliases = slices.Collect(maps.Keys(aliasSet))

	commandSet := make(map[string]bool)
	for _, bin := range explorePath(diff.Path) {
		commandSet[bin] = true
	}
	relations.Commands = slices.Collect(maps.Keys(commandSet))

	return
}

func explorePath(path []string) (bins []string) {
	bins = make([]string, 0)
	for _, entry := range path {
		dirEntries, err := os.ReadDir(entry)
		if err != nil {
			continue
		}
		for _, dirEntry := range dirEntries {
			if dirEntry.IsDir() {
				continue
			}
			bins = append(bins, dirEntry.Name())
		}
	}
	return
}

func compareEnvs(before, after Env) (difference Env) {
	difference.Aliases = make(map[string]string)
	for key, value := range after.Aliases {
		beforeValue, ok := before.Aliases[key]
		if value != beforeValue || !ok {
			difference.Aliases[key] = value
		}
	}
	difference.Functions = make(map[string]string)
	for key, value := range after.Functions {
		beforeValue, ok := before.Functions[key]
		if value != beforeValue || !ok {
			difference.Functions[key] = value
		}
	}
	difference.Path = make([]string, 0)
	for _, entry := range after.Path {
		if !slices.Contains(before.Path, entry) {
			difference.Path = append(difference.Path, entry)
		}
	}
	return
}

func loadEnv(prefix string) (env Env) {
	env.Aliases = loadAliases(prefix)
	env.Functions = loadFunctions(prefix)
	env.Path = loadPath(prefix)
	return
}

func loadAliases(prefix string) (aliases map[string]string) {
	aliases = make(map[string]string)
	for line := range strings.SplitSeq(shellCmd(prefix + "alias"), "\n") {
		if line == "" {
			continue
		}
		groups := aliasRegex.FindStringSubmatch(line)
		if len(groups) < 3 {
			continue
		}
		aliases[groups[1]] = groups[2]
	}
	return
}

func loadFunctions(prefix string) (functions map[string]string) {
	functions = make(map[string]string)
	cmdParts := make([]string, 0)
	functionNames := strings.Split(shellCmd(prefix + "functions"), "\n")
	for _, name := range functionNames {
		cmdParts = append(cmdParts, "functions --details " + name)
	}
	for idx, detail := range strings.Split(shellCmd(prefix + strings.Join(cmdParts, " && ")), "\n") {
		functions[functionNames[idx]] = detail
	}
	return
}

func loadPath(prefix string) (path []string) {
	path = make([]string, 0)
	for entry := range strings.SplitSeq(shellCmd(prefix + "echo $PATH"), " ") {
		if entry == "." {
			continue
		}
		path = append(path, entry)
	}
	return
}
