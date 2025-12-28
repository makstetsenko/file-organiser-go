package args

import "flag"

type AppArgs struct {
	ConfigFilePath string
}

func GetArgs() AppArgs {
	configPath := flag.String("config", "config.yaml", "Path to file with organization configuration. Check config.example.yaml")

	flag.Parse()

	return AppArgs{
		ConfigFilePath: *configPath,
	}
}
