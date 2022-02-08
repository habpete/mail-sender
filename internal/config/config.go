package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func Read(refConfig interface{}, pathToConfig string) error {
	data, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("read config %s failed %w", pathToConfig, err)
	}

	if err = json.Unmarshal(data, refConfig); err != nil {
		return fmt.Errorf("unmarshal config %s failed %w", pathToConfig, err)
	}

	return nil
}

type SmtpConfig struct {
	RemoteHost        string        `json:"remote_host"`
	User              string        `json:"user"`
	From              string        `json:"from"`
	ServiceName       string        `json:"service_name"`
	Password          string        `json:"password"`
	ServePort         string        `json:"serve_port"`
	ReconnectInterval time.Duration `json:"reconnect_interval"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"database_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
