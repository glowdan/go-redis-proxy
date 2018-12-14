package main

import (
	redisServer "github.com/dotcloud/go-redis-server"
	"os"
)

func main() {

	server, err := redisServer.NewServer(DefaultConfig())
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe(); err != nil {
		_ = os.Remove("/tmp/redis.sock")
		panic(err)
	}
}
