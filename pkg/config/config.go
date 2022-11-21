package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Server Server `yaml:"server"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

func ReadConfig() {
	byt, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalf("Failed to ready file with err : %v", err)
	}

	err = yaml.Unmarshal(byt, &Cfg)
	if err != nil {
		log.Fatalf("Failed to parse config into struct with err : %v", err)
	}
}
