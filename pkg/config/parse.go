package config

import (
	"encoding/json"
	"os"

	"github.com/rusik69/servicequotas/pkg/types"
)

// Parse parses the quotas configuration file
func Parse(fileName string) (types.QuotasConfig, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return types.QuotasConfig{}, err
	}
	defer file.Close()
	var cfg types.QuotasConfig
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return types.QuotasConfig{}, err
	}
	return cfg, nil
}
