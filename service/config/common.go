package config

import (
	"golang.org/x/time/rate"
)

type Config struct {
	Env string `yaml:"env"`
	Db  struct {
		Type         string `yaml:"type"`
		MaxIdleConns int    `yaml:"max-idle-conns"`
		MaxOpenConns int    `yaml:"max-open-conns"`
		Port         string `yaml:"port"`
		Host         string `yaml:"host"`
		TablePrefix  string `yaml:"table_prefix"`
		Database     string `yaml:"database"`
		User         string `yaml:"name"`
		PassWord     string `yaml:"password"`
		Source       string `yaml:"source"`
	}
	MyJwt struct {
		Secret    string `yaml:"secret"`
		ExpiresAt int64  `yaml:"expires_at"`
	}
	App struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		UploadFile string `yaml:"uploadFile"`
	}
	Rate struct {
		Limit rate.Limit `yaml:"limit"`
		Burst int        `yaml:"burst"`
	}
	Logger struct {
		Drive  string `yaml:"drive"`
		Path   string `yaml:"path"`
		Size   int    `yaml:"size"`
		MaxAge int    `yaml:"maxAge"`
		StdOut bool   `yaml:"stdOut"`
	}
}
