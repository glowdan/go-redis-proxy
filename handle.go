package main

import (
	"encoding/json"
	"fmt"
	redisServer "github.com/dotcloud/go-redis-server"
)

type MyHandler struct {
	redisServer.DefaultHandler
}

//Request {"id":"4","method":"1","action":"2","args":"3"}
// Get override the DefaultHandler's method.
func (h *MyHandler) Get(key string) ([]byte, error) {
	var request Request
	err := json.Unmarshal([]byte(key), &request)
	if err != nil {
		fmt.Println("error:", err)
	}
	response, err := dispacher(request)
	if err != nil {
		return nil, err
	}
	return json.Marshal(response)
}
