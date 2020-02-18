package configs

import (
	"fmt"
	"os"
)

type MongoDB struct {
	Host string
	Port string
	DB   string
}

func GetMongoDB() *MongoDB {
	return &MongoDB{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DB:   os.Getenv("DB_DATABASE"),
	}
}

func (config *MongoDB) GetConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)
}
