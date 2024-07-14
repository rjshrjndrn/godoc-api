package config

import (
	"context"
	"fmt"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type EnvConfig struct {
	Database *DatabaseConfig
}

type DatabaseConfig struct {
	Port      string `env:"DBPORT"`
	Username  string `env:"DBUSERNAME"`
	Password  string `env:"DBPASSWORD"`
	Host      string `env:"DBHOST"`
	Name      string `env:"DBNAME"`
	DBConnUrl string
}

// parse configs
func ParseConfig() (*EnvConfig, error) {
	var env EnvConfig
	err := envconfig.Process(context.Background(), &env)
	if err != nil {
		return nil, err
	}
	env.Database.DBConnUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", env.Database.Username, env.Database.Password, env.Database.Host, env.Database.Port, env.Database.Name)
	log.Print(env.Database.DBConnUrl)
	return &env, nil
}
