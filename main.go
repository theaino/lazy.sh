package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)


func main() {
	loadCache()
	defer fmt.Println(startPath)

	input, err := io.ReadAll(os.Stdin)
	handle(err)

	matches, sum := compareSum(input, sumPath)
	if matches {
		return
	}

	createScript(string(input))
	writeSum(sum)
}

func createScript(input string) {
	log.Print("Caching the initializer relations. This might take some time.")

	startFile, err := os.OpenFile(startPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	defer startFile.Close()
	handle(err)

	loaders := make([]Loader, 0)
	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}
		loaders = append(loaders, fetchLoader(line))
	}
	_, err = startFile.WriteString(formatLoaders(loaders))
	handle(err)
}

