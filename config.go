package main

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Configure struct {
	Debug bool

	Server struct {
		Base string
		Port int
	}

	Database struct {
		Host     string
		User     string
		Password string
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
	conf.Debug = true
	conf.Server.Base = "http://localhost:8080"
	conf.Server.Port = 8080
	conf.Database.Host = "localhost"
	conf.Database.User = "media_expose"
	conf.Database.Password = "change_this_password"

	enc := yaml.NewEncoder(file)
	defer enc.Close()

	if err = enc.Encode(conf); err != nil {
		panic(err)
	}
	return nil
}
