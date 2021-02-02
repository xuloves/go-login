package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App   `yaml:"app"`
	Sms   `yaml:"sms"`
	Mysql `yaml:"mysql"`
	Redis `yaml:"redis"`
}

type App struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type Sms struct {
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	RegionId        string `yaml:"regionId"`
	TemplateCode    string `yaml:"templateCode"`
	SignName        string `yaml:"signName"`
}
type Mysql struct {
	Path     string `yaml:"path"`
	Config   string `yaml:"config"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"DB"`
}

var config *Config = nil

func GetConfig() *Config {
	return config

}
func ParseConfig(path string) (*Config, error) {
	config = &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(config)
	}
	return config, nil
}
