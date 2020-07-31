package main

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Configure struct {
	Mode string

	Server struct {
		Base string
		Port int
	}

	Database struct {
		Host     string
		User     string
		Password string

		Path string // used for specify sqlite path
	}
}

func LoadConfigure(path string) (conf Configure, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	dec := yaml.NewDecoder(file)
	err = dec.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return
}

func InitialConfigure(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	var conf Configure
	conf.Mode = "sqlite"
	conf.Server.Base = "http://localhost:8080"
	conf.Server.Port = 8080
	conf.Database.Path = "gorm.db"

	enc := yaml.NewEncoder(file)
	defer enc.Close()

	if err = enc.Encode(conf); err != nil {
		panic(err)
	}
	return nil
}
