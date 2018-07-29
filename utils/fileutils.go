package utils

import (
	"os"
	"path/filepath"
	"sync"
)

const (
	ConfFile = "conf.yaml"
	ConfDir  = "conf"
)

var (
	once    sync.Once
	confDir string
)

func GetWorkDir() string {
	once.Do(func() {
		wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			confDir = "./"
		}
		confDir = wd
	})
	return confDir
}

func GetConfDir() string {
	return filepath.Join(GetWorkDir(), ConfDir)
}

func GetConfFile() string {
	return filepath.Join(GetConfDir(), ConfFile)
}
