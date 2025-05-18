package main

import (
	"crypto/md5"
	"io"
	"os"
	"path/filepath"
)


var cacheRoot string
var startPath string
var sumPath string

func loadCache() {
	cacheDir, err := os.UserCacheDir()
	handle(err)
	cacheRoot = filepath.Join(cacheDir, "lazysh")
	os.MkdirAll(cacheRoot, os.ModePerm)

	startPath = filepath.Join(cacheRoot, "start.fish")
	sumPath = filepath.Join(cacheRoot, "stdin.sum")
}

func compareSum(input []byte, sumPath string) (bool, []byte) {
	sumFile, err := os.OpenFile(sumPath, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	defer sumFile.Close()

	rawSum := md5.Sum(input)
	sum := rawSum[:]

	oldSum, err := io.ReadAll(sumFile)
	handle(err)

	matches := string(sum) == string(oldSum)
	return matches, sum
}

func writeSum(sum []byte) {
	sumFile, err := os.OpenFile(sumPath, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	handle(err)
	defer sumFile.Close()

	_, err = sumFile.Write(sum)
	handle(err)
}
