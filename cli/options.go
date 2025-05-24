package cli

import (
	"flag"
	"lazysh/shell"
)

type Options struct {
	ForceAnalyze bool
	Shell        shell.Shell
}

func ParseOptions() (options Options, err error) {
	flag.BoolVar(&options.ForceAnalyze, "f", false, "force to analyze init commands")
	flag.Parse()

	flag.Parse()

	var shellName string
	if flag.NArg() > 0 {
		shellName = flag.Arg(0)
	} else {
		shellName, err = fetchShell()
		if err != nil {
			return
		}
	}
	options.setShell(shellName)

	return
}

func (o *Options) setShell(name string) {
	switch name {
	case "bash":
		o.Shell = shell.Bash{}
	case "zsh":
		o.Shell = shell.Zsh{}
	case "fish":
		o.Shell = shell.Fish{}
	default:
		o.Shell = shell.Bash{}
	}
}
