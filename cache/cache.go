package cache

import (
	"bytes"
	"io"
	"lazysh/shell"
	"os"
	"path/filepath"
)

type Cache struct {
	RootDir    string
	ScriptPath string
	SumPath    string
}

func LoadCache(s shell.Shell) (cache Cache, err error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return
	}
	cache.RootDir = filepath.Join(cacheDir, "lazysh")
	os.MkdirAll(cache.RootDir, os.ModePerm)

	cache.ScriptPath = filepath.Join(cache.RootDir, "start"+s.Extension())
	cache.SumPath = filepath.Join(cache.RootDir, "start" + s.Extension() + ".sum")
	return
}

func (c Cache) CheckSum(sum []byte) (matches bool, err error) {
	sumFile, err := os.OpenFile(c.SumPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return false, nil
	}
	defer sumFile.Close()

	oldSum, err := io.ReadAll(sumFile)
	if err != nil {
		return
	}

	matches = bytes.Compare(oldSum, sum) == 0
	return
}

func (c Cache) WriteScript(data string) (err error) {
	scriptFile, err := os.OpenFile(c.ScriptPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer scriptFile.Close()

	_, err = scriptFile.WriteString(data)
	return
}

func (c Cache) WriteSum(sum []byte) (err error) {
	sumFile, err := os.OpenFile(c.SumPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer sumFile.Close()

	_, err = sumFile.Write(sum)
	return
}
