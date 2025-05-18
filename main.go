package main

import (
	"lazysh/cli"
	"log"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := cli.NewCli()
	handle(c.Load())
	handle(c.Run())
}
