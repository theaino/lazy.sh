package shell

type Shell interface {
	Aliases(prefix string) (map[string]string, error)
	Functions(prefix string) (map[string]string, error)
	Path(prefix string) ([]string, error)
	cmd(cmd string) (string, error)

	FormatAlias(name string, cmd string) string
	FormatCommand(name string, cmd string) string

	Extension() string
}
