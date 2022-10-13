package config

import (
	"os"

	"github.com/naoina/toml"
)

type DataBase struct {
	DataSource string
	UserName   string
	PassWord   string
	DB         string
	Collection string
}

type Config struct {
	DataBase
}

func NewConfig(file string) *Config {
	c := new(Config)

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer f.Close()
		if err := toml.NewDecoder(f).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
