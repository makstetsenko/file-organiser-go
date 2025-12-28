package main

import (
	"github.com/makstetsenko/file-organiser-go/args"
	"github.com/makstetsenko/file-organiser-go/domain"
)

func main() {
	appArgs := args.GetArgs()

	conf := domain.ReadConfig(appArgs.ConfigFilePath)

	domain.MoveFiles(conf)
}
