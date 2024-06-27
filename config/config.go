package config

import "os"

type App struct {
	Host string
	Port string
}

func NewAppConf() *App {
	return &App{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}
}

type DB struct {
	Driver   string 
	Name     string 
	User     string 
	Password string 
	Host     string
	Port     string 
}

func NewDBConf() DB {
	return DB{
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
