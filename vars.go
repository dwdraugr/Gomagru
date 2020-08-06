package main

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis"
)

type dbConf struct {
	Host string `json:"host"`
	Name string `json:"name"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type redisConf struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Pass string `json:"pass"`
}

type config struct {
	DbConf    dbConf    `json:"db"`
	RedisConf redisConf `json:"redis"`
}

var (
	Config  config
	Context context.Context
	Db      *sql.DB
	Redis   *redis.Client
)
