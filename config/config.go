package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Database Database
}

type Database struct {
	Mysql    Mysql
	Postgres Postgres
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
}

// singleton pattern

var (
	// cfg is the singleton configuration instance
	cfg Config
	// once is used for the singleton pattern
	once sync.Once
)

func envString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func envInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		// convert string to int
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Sprintf("Invalid value for %s: %s", key, value))
		}

		return valueInt
	}
	return fallback
}

// return the configuration

func Get() Config {
	once.Do(func() {
		// load configuration from environment variables
		godotenv.Load()

		cfg = Config{
			Database: Database{
				Mysql: Mysql{
					Host:     envString("MYSQL_HOST", "localhost"),
					Port:     envInt("MYSQL_PORT", 3306),
					User:     envString("MYSQL_USER", "root"),
					Password: envString("MYSQL_PASSWORD", "password"),
					Database: envString("MYSQL_DATABASE", "mydb"),
				},
				Postgresql: Postgresql{
					Host:     envString("POSTGRES_HOST", "localhost"),
					Port:     envInt("POSTGRESQL_PORT", 5432),
					Username: envString("POSTGRESQL_USERNAME", "postgres"),
					Password: envString("POSTGRESQL_PASSWORD", "password"),
					Database: envString("POSTGRESQL_DATABASE", "mydb"),
				},
			},
		}
	})
	return cfg
}
