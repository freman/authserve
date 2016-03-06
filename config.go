package main

import (
	t "github.com/freman/go-commontypes"
)

type jwtConfig struct {
	Private   t.KeyFile
	Public    t.KeyFile
	Expiry    t.Duration
	NotBefore t.Duration `toml:"not_before"`
}

type databaseConfig struct {
	Driver string
	DSN    string
}

type configuration struct {
	JWT      jwtConfig
	Database databaseConfig
}
