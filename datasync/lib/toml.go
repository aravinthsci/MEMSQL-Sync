package lib

import (
	. "datasync/constants"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/k0kubun/pp"
)

type tomlConfig struct {
	Database map[string]Database
	SSH      map[string]SSH
}

func LoadTomlConf(configPath string) (tmlconf tomlConfig) {
	log.Print("[Setting] loading toml configuration...")
	if _, err := toml.DecodeFile(configPath, &tmlconf); err != nil {
		pp.Print(err)
	}

	log.Print("[Setting] loaded toml configuration")
	return tmlconf
}
