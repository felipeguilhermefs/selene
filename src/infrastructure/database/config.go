package database

import "time"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Conn     ConnConfig
}

type ConnConfig struct {
	MaxIdle int
	MaxOpen int
	TTL     time.Duration
}
