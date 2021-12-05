package main

import (
  "log"

	"github.com/go-redis/redis/v8"
)


var rds *redis.Client


func init_redis() {

	rds = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

} // init_redis


func main() {
	
	init_redis()
	log.Printf("Starting %s service...", version())	
	
} // main
