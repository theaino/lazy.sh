package shell

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Fish struct{}

func (f Fish) cmd(cmd string) (string, error) {
	command := exec.Command("fish", "-Nc", cmd)
	command.Env = append(command.Env, "DISABLE_LAZY=1")
	out, err := command.Output()
	if err != nil {
		return "", fmt.Errorf("Failed to run %s: %w", cmd, err)
	}
	return strings.Trim(string(out), "\n"), nil
}

func (f Fish) Aliases(prefix string) (aliases map[string]string, err error) {
	aliasRegex := regexp.MustCompile(`alias ([^\ ]+) ([^\ ]+)`)

	rawAliases, err := f.cmd(prefix + "alias")
	if err != nil {
		return
	}

	aliases = make(map[string]string)
	for line := range strings.SplitSeq(rawAliases, "\n") {
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

func (f Fish) Functions(prefix string) (functions map[string]string, err error) {
	names, err := f.functionNames(prefix)
	if err != nil {
		return
	}

	return f.functionDetails(prefix, names)
}

func (f Fish) functionNames(prefix string) (names []string, err error) {
	rawNames, err := f.cmd(prefix + "functions")
	if err != nil {
		return
	}
	names = strings.Split(rawNames, "\n")
	return
}

func (f Fish) functionDetails(prefix string, names []string) (details map[string]string, err error) {
	cmd := make([]string, len(names))
	for idx, name := range names {
		cmd[idx] = "functions --details " + name
	}

	rawDetails, err := f.cmd(prefix + strings.Join(cmd, " && "))
	if err != nil {
		return
	}

	details = make(map[string]string)
	for idx, detail := range strings.Split(rawDetails, "\n") {
		details[names[idx]] = detail
	}
	return
}

func (f Fish) Path(prefix string) (entries []string, err error) {
	rawPath, err := f.cmd(prefix + "echo $PATH")
	if err != nil {
		return
	}

	entries = strings.Split(rawPath, " ")
	return
}
