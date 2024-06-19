package config

import (
	"encoding/json"
	"github.com/oduortoni/odrserver/src/server/entities"
	"os"
)

func Read(filename string) (bool, *entities.Config) {
	configFD, err := os.Open(filename)
	if err != nil {
		return false, nil
	}

	config := entities.Config{}
	json.NewDecoder(configFD).Decode(&config)
	return true, &config
}
