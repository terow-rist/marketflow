package config

import (
	"os"
	"strconv"
)

type (
	App struct {
		Info  *Info
		DB    *DB
		Redis *Redis
	}

	Info struct {
		Name string
		Env  string
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		Database int
	}
)

func Init() *App {
	info := &Info{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	dbNum, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))

	rdb := &Redis{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: dbNum,
	}

	return &App{
		Info:  info,
		DB:    db,
		Redis: rdb,
	}
}

func (r *Redis) Addr() string {
	return r.Host + ":" + r.Port
}
