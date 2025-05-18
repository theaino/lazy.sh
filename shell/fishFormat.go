package shell

import (
	"fmt"
	"strings"
)

func (f Fish) FormatAlias(name string, cmd string, aliasFunction string) string {
	return fmt.Sprintf(`function %s;%s;%s;%s $argv;end`, name, aliasFunction, cmd, name)
}

func (f Fish) FormatCommand(name string, cmd string, aliasFunction string) string {
	return fmt.Sprintf(`function %s;%s;%s;%s $argv;end`, name, aliasFunction, cmd, name)
}

func (f Fish) FormatCommandAliasFunction(name string, aliases []string) string {
	formattedAliases := make([]string, len(aliases))
	for idx, alias := range aliases {
		formattedAliases[idx] = fmt.Sprintf(`alias %s="command %s"`, alias, alias)
	}
	return fmt.Sprintf(`function %s;%s;end`, name, strings.Join(formattedAliases, ";"))
}

func (f Fish) Extension() string {
	return ".fish"
}
