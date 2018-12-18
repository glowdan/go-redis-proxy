package main

import (
	"errors"
	"log"
)

func dispacher(request Request) (Response, error) {
	var response = NewResponse()
	var data interface{}
	var err error
	ec := 0
	switch request.Method {
	case "redis":
		data, err = RedisAction(request.Action, request.Args)
	case "cache":
		data, err = CacheAction(request.Action, request.Args)
	default:
		err = errors.New("Unsupported method ")
		data = ""
		ec = 405
	}
	if err != nil {
		log.Printf("id: %s,error: %v", request.Id, err)
		response.Em = err.Error()
		if ec == 0 {
			ec = 400
		}
		response.Ec = ec
	}
	response.Data = data
	return response, nil
}
