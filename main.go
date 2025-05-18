package main

import (
	"flag"
	"fmt"
	"io"
	"lazysh/cache"
	"lazysh/load"
	"lazysh/shell"
	"log"
	"os"
	"strings"
)

func main() {
	s := getShell()
	c, err := cache.LoadCache(s)
	if err != nil {
		log.Panic(err)
	}
	defer fmt.Println(c.ScriptPath)

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Panic(err)
	}

	sum := cache.Sum(input)
	matches, err := c.CheckSum(sum)
	if err != nil {
		log.Panic(err)
	}
	if matches {
		return
	}

	data, err := createScript(s, string(input))
	if err != nil {
		log.Panic(err)
	}
	if err := c.WriteScript(data); err != nil {
		log.Panic(err)
	}
	if err := c.WriteSum(sum); err != nil {
		log.Panic(err)
	}
}

func getShell() shell.Shell {
	flag.Parse()
	switch flag.Arg(0) {
	case "bash":
		return shell.Bash{}
	case "fish":
		return shell.Fish{}
	}
	return shell.Bash{}
}

func createScript(s shell.Shell, input string) (data string, err error) {
	log.Print("Caching the initializer relations. This might take some time.")

	loaders := make([]load.Loader, 0)
	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}
		loaders = append(loaders, load.Loader{})
		loaders[len(loaders)-1], err = load.NewLoader(s, line)
		if err != nil {
			return
		}
	}
	data = load.FormatLoaders(s, loaders)
	return
}
