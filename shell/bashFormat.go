package shell

import "fmt"

func (b Bash) FormatAlias(name string, cmd string) string {
	return fmt.Sprintf(`%s() { %s;%s $@; }`, name, cmd, name)
}

func (b Bash) FormatCommand(name string, cmd string) string {
	return fmt.Sprintf(`%s() { %s;alias %s="command %s";%s $@; }`, name, cmd, name, name, name)
}

func (b Bash) Extension() string {
	return ".bash"
}
