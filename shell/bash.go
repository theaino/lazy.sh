package shell

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Bash struct{}

func (b Bash) cmd(cmd string) (string, error) {
	command := exec.Command("bash", "--norc", "-c", cmd)
	command.Env = os.Environ()
	command.Env = append(command.Env, "DISABLE_LAZY=1")
	out, err := command.Output()
	if err != nil {
		return "", fmt.Errorf("Failed to run %s: %w", cmd, err)
	}
	return strings.Trim(string(out), "\n"), nil
}

func (b Bash) MakePrefix(cmd string) string {
	return cmd + " 1>&2 && "
}

func (b Bash) Aliases(prefix string) (aliases map[string]string, err error) {
	aliasRegex := regexp.MustCompile(`alias ([^\ ]+)='(.+)'`)

	rawAliases, err := b.cmd(prefix + "alias")
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

func (b Bash) Functions(prefix string) (functions map[string]string, err error) {
	functionRegex := regexp.MustCompile(`([^\ \n]+)\ *\(\)\ *\n{\ *((?:\n[\ \t]+.*)+)\n}`)

	rawFunctions, err := b.cmd(prefix + "declare -f")
	if err != nil {
		return
	}
	
	functions = make(map[string]string)
	for _, groups := range functionRegex.FindAllStringSubmatch(rawFunctions, -1) {
		if len(groups) < 3 {
			continue
		}
		functions[groups[1]] = groups[2]
	}
	return
}

func (b Bash) Path(prefix string) (entries []string, err error) {
	rawPath, err := b.cmd(prefix + "echo $PATH")
	if err != nil {
		return
	}

	entries = strings.Split(rawPath, ":")
	return
}
