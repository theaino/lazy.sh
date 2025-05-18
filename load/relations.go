package load

import (
	"lazysh/shell"
	"maps"
	"os"
	"slices"
)

type Env struct {
	Aliases   map[string]string
	Functions map[string]string
	Path      []string
}

type Relations struct {
	Aliases  []string
	Commands []string
}

func Analyze(s shell.Shell, cmd string) (relations Relations, err error) {
	prefix := s.MakePrefix(cmd)
	before, err := loadEnv(s, "")
	if err != nil {
		return
	}
	after, err := loadEnv(s, prefix)
	if err != nil {
		return
	}
	difference := compareEnvs(before, after)

	aliasSet := make(map[string]bool)
	for alias := range difference.Aliases {
		aliasSet[alias] = true
	}
	for function := range difference.Functions {
		aliasSet[function] = true
	}
	relations.Aliases = slices.Collect(maps.Keys(aliasSet))

	commandSet := make(map[string]bool)
	for _, bin := range explorePath(difference.Path) {
		commandSet[bin] = true
	}
	relations.Commands = slices.Collect(maps.Keys(commandSet))
	return
}

func explorePath(path []string) (bins []string) {
	bins = make([]string, 0)
	for _, entry := range path {
		if entry == "." {
			continue
		}
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

func loadEnv(s shell.Shell, prefix string) (env Env, err error) {
	env.Aliases, err = s.Aliases(prefix)
	if err != nil {
		return
	}
	env.Functions, err = s.Functions(prefix)
	if err != nil {
		return
	}
	env.Path, err = s.Path(prefix)
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
