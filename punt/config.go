package punt

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Clusters map[string]ClusterConfig `json:"clusters"`
	Indexes  []IndexConfig            `json:"indexes"`
	Mappings []Mapping                `json:"mappings"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = json.Unmarshal(file, &config)
	return &config, nil
}
