package config

import "github.com/BurntSushi/toml"

type Config struct {
	App *APP
	DB  *DB
}
type APP struct {
	Mode string
}

type DB struct {
	Type string
	DSN  string
}

var Conf Config

func Setup() {
	if _, err := toml.DecodeFile("./config/conf.toml", &Conf); err != nil {
		panic(err)
	}
}
