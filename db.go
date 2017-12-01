package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

var samples = []check{
	check{Name: "the googs", Address: "http://google.com"},
	check{Name: "zbos address book", Address: "http://facebook.com"},
}

func initializeDB() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func seedDB() {
	expiry := time.Hour * 24
	for _, s := range samples {
		redisClient.Set(s.Name, s.Address, expiry)
	}

	printDBSize()
}

func dropDB() {
	redisClient.FlushAll()
	printDBSize()
}

func printDBSize() {
	fmt.Println("entries:", redisClient.DbSize())
}
