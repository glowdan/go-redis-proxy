package main

import (
	redisServer "github.com/dotcloud/go-redis-server"
)

func main() {

	myhandler := &MyHandler{}

	server, err := redisServer.NewServer(redisServer.DefaultConfig().Handler(myhandler))
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
