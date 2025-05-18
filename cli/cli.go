package cli

import (
	"fmt"
	"io"
	"lazysh/cache"
	"lazysh/load"
	"log"
	"os"
	"strings"
)

type Cli struct {
	Options Options
	Cache   cache.Cache
	Sum     []byte
	Input   string
}

func NewCli() Cli {
	return Cli{}
}

func (c *Cli) Load() (err error) {
	c.Options = ParseOptions()
	c.Cache, err = cache.LoadCache(c.Options.Shell)
	if err != nil {
		return
	}
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}
	c.Input = string(input)
	c.Sum = cache.Sum(input)
	return
}

func (c *Cli) Run() (err error) {
	defer fmt.Println(c.Cache.ScriptPath)

	matches, err := c.Cache.CheckSum(c.Sum)
	if err != nil {
		return
	}
	if c.Options.ForceAnalyze || !matches {
		err = c.analyze()
	}
	return
}

func (c *Cli) analyze() (err error) {
	log.Println("Analyzing. This might take some time.")

	err = c.generateScript()
	if err != nil {
		return
	}
	err = c.Cache.WriteSum(c.Sum)
	return
}

func (c *Cli) generateScript() (err error) {
	loaders := make([]load.Loader, 0)
	for line := range strings.SplitSeq(c.Input, "\n") {
		if line == "" {
			continue
		}
		loaders = append(loaders, load.Loader{})
		loaders[len(loaders)-1], err = load.NewLoader(c.Options.Shell, line)
		if err != nil {
			return
		}
	}
	err = c.Cache.WriteScript(load.FormatLoaders(c.Options.Shell, loaders))
	return
}
