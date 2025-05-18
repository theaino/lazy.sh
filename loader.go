package main

import (
	"fmt"
	"strings"
)


type Loader struct {
	InitCmd string
	Relations Relations
}

func fetchLoader(cmd string) Loader {
	return Loader{
		InitCmd: cmd,
		Relations: loadRelations(cmd),
	}
}

func formatLoaders(loaders []Loader) string {
	loaderStrings := make([]string, len(loaders))
	for idx, loader := range loaders {
		loaderStrings[idx] = formatLoader(loader)
	}
	return strings.Join(loaderStrings, "\n")
}

func formatLoader(loader Loader) string {
	lines := make([]string, 0)
	for _, alias := range loader.Relations.Aliases {
		lines = append(lines, formatAlias(loader.InitCmd, alias))
	}
	for _, command := range loader.Relations.Commands {
		lines = append(lines, formatCommand(loader.InitCmd, command))
	}
	return strings.Join(lines, "\n")
}

func formatAlias(cmd string, name string) string {
	return fmt.Sprintf(`function %s;%s;%s $argv;end`, name, cmd, name)
}

func formatCommand(cmd string, name string) string {
	return fmt.Sprintf(`function %s;%s;alias %s="command %s";%s $argv;end`, name, cmd, name, name, name)
}
