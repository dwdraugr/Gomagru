package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
)

func Init() {
	readConfig()
	initDb()
	initRedis()
}

func readConfig() {
	rawFile, err := os.Open("config.json")
	if err != nil {
		panic("Ooops, config.json is note load. Message:\n" + err.Error())
	}
	jsonFile, _ := ioutil.ReadAll(rawFile)
	err = json.Unmarshal(jsonFile, &Config)
	if err != nil {
		panic("Ooops, config.json in bad. Message:\n" + err.Error())
	}
	return
}

func initDb() {
	var err error
	Db, err = sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			Config.DbConf.User,
			Config.DbConf.Pass,
			Config.DbConf.Host,
			Config.DbConf.Port,
			Config.DbConf.Name,
		),
	)
	if err != nil {
		panic("Ooops, trouble with connect to db. Message:\n" + err.Error())
	}
}

func initRedis() {
	Context = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			Config.RedisConf.Host,
			Config.RedisConf.Port),
		Password: Config.RedisConf.Pass,
		DB:       0,
	})
	_, err := rdb.Ping(Context).Result()
	if err != nil {
		panic("Ooops, touble with connect ti redis. Message:\n" + err.Error())
	}
}
