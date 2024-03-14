package config

import "vec-calc-server/lib"

type Config struct {
	lib.ServerConfig `json:"-"`
	lib.DatabaseConfig
}
