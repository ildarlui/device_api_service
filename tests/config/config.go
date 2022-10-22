package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Host      string `envconfig:"HOST" default:"127.0.0.1:8000"`
	Livecheck string `envconfig:"LIVECHECK" default:"/live"`
	BaseUrl   string `envconfig:"BASE_URL" default:""`
}

func GetConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Config process error: %v", err)
	}
	return config, err
}

//type Config2 struct {
//	Host string `yaml:"base_url"`
//	Livecheck string
//}
//
//func NewFromYaml(filePath string) (*Config2, error) {
//	cfg := &Config2{}
//	file, err := os.Open(filepath.Clean(filePath))
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	decoder := yaml.NewDecoder(file)
//	if err := decoder.Decode(cfg); err != nil {
//		return nil, err
//	}
//	return cfg, nil
//}
