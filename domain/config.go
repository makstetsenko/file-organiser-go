package domain

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type OrganisationConfig struct {
	FileMovingConfigs []FileOrganisationConfig `yaml:"files_organisation"`
}

type FileOrganisationConfig struct {
	SourceFiles []string `yaml:"source"`
	Destination string   `yaml:"destination"`
}

func ReadConfig(path string) OrganisationConfig {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Cannot open file by path %v.\n%v\n", path, err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var conf OrganisationConfig

	if err := decoder.Decode(&conf); err != nil {
		log.Fatalf("Cannot read config %v.\n%v\n", path, err)
	}

	return conf
}
