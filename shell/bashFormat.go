package shell

import (
	"fmt"
	"strings"
)

func (b Bash) FormatAlias(name string, cmd string, aliasFunction string) string {
	return fmt.Sprintf(`%s() { %s;%s;%s $@; }`, name, aliasFunction, cmd, name)
}

func (b Bash) FormatCommand(name string, cmd string, aliasFunction string) string {
	return fmt.Sprintf(`%s() { %s;%s;%s $@; }`, name, aliasFunction, cmd, name)
}

func (b Bash) FormatCommandAliasFunction(name string, aliases []string) string {
	formattedAliases := make([]string, len(aliases))
	for idx, alias := range aliases {
		formattedAliases[idx] = fmt.Sprintf(`%s () { command %s $@; }`, alias, alias)
	}
	aliasList := strings.Join(formattedAliases, ";")
	if aliasList == "" {
		aliasList = ":"
	}
	return fmt.Sprintf(`%s() { %s; }`, name, aliasList)
}

func (b Bash) Extension() string {
	return ".bash"
}
