package shell

type Shell interface {
	MakePrefix(cmd string) string
	Aliases(prefix string) (map[string]string, error)
	Functions(prefix string) (map[string]string, error)
	Path(prefix string) ([]string, error)
	cmd(cmd string) (string, error)

	FormatAlias(name string, cmd string, aliasFunction string) string
	FormatCommand(name string, cmd string, aliasFunction string) string
	FormatCommandAliasFunction(name string, aliases []string) string

	Extension() string
}
