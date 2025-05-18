package shell

import "fmt"

func (f Fish) FormatAlias(name string, cmd string) string {
	return fmt.Sprintf(`function %s;%s;%s $argv;end`, name, cmd, name)
}

func (f Fish) FormatCommand(name string, cmd string) string {
	return fmt.Sprintf(`function %s;%s;alias %s="command %s";%s $argv;end`, name, cmd, name, name, name)
}

func (f Fish) Extension() string {
	return ".fish"
}
