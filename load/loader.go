package load

import (
	"fmt"
	"lazysh/shell"
	"strings"
)

type Loader struct {
	InitCmd   string
	Relations Relations
}

func NewLoader(s shell.Shell, cmd string) (loader Loader, err error) {
	loader.Relations, err = Analyze(s, cmd)
	if err != nil {
		return
	}
	loader.InitCmd = cmd
	return
}

func FormatLoaders(s shell.Shell, loaders []Loader) string {
	loaderStrings := make([]string, len(loaders))
	for idx, loader := range loaders {
		loaderStrings[idx] = formatLoader(s, loader, idx)
	}
	return strings.Join(loaderStrings, "\n")
}

func formatLoader(s shell.Shell, loader Loader, idx int) string {
	lines := make([]string, 1)
	commandAliasFunctionName := fmt.Sprintf("__lazysh_command_alias_%d", idx)
	lines[0] = s.FormatCommandAliasFunction(commandAliasFunctionName, append(loader.Relations.Aliases, loader.Relations.Commands...))

	for _, alias := range loader.Relations.Aliases {
		lines = append(lines, s.FormatAlias(alias, loader.InitCmd, commandAliasFunctionName))
	}
	for _, command := range loader.Relations.Commands {
		lines = append(lines, s.FormatCommand(command, loader.InitCmd, commandAliasFunctionName))
	}
	return strings.Join(lines, "\n")
}
