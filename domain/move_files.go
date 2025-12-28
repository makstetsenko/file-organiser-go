package domain

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func MoveFiles(conf OrganisationConfig) {
	counter := 0
	failedCounter := 0
	for _, fileConfig := range conf.FileMovingConfigs {
		var files []string

		for _, p := range fileConfig.SourceFiles {
			configPathList, err := filepath.Glob(p)

			if err != nil {
				log.Fatalf("Cannot get list of files %v.\n%v\n", p, err)
			}

			files = append(files, configPathList...)
		}

		if err := os.MkdirAll(fileConfig.Destination, 0755); err != nil {
			log.Fatalf("Cannot create reports directory %v", fileConfig.Destination)
		}

		for _, oldFilePath := range files {
			newFilePath := fmt.Sprintf("%v/%v", fileConfig.Destination, path.Base(oldFilePath))

			if err := os.Rename(oldFilePath, newFilePath); err != nil {
				log.Printf("Cannot move file from %v to %v.\n%v\n", oldFilePath, newFilePath, err)
				failedCounter++
				continue
			}

			log.Printf("Moved file from %v to %v.\n", oldFilePath, newFilePath)
			counter++
		}
	}
	log.Printf("Moved %v files.\n", counter)

	if failedCounter > 0 {
		log.Printf("Failed to move %v files.\n", failedCounter)
	}
}
