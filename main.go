package main

import (
  "context"
	"flag"
	"fmt"
  //"log"

	"github.com/go-redis/redis/v8"
)


var rds *redis.Client

var ctx = context.Background()


var (
	host 		= flag.String("host", DEFAULT_REDIS_HOST, "redis host address")
	port    = flag.String("port", DEFAULT_REDIS_PORT, "redis port]")
)


func init_redis() {

	rds = redis.NewClient(&redis.Options{
		Addr: addr(),
		Password: "",
		DB: 0,
	})

} // init_redis


func main() {
	
	flag.Parse()

	init_redis()
	fmt.Printf("Starting %s service...\n", version())	

	registerUser("a@aol.com", "test", "amuncher")

	createCharacter("test", false)
	
} // main
