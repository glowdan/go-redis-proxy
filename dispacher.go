package main

import "fmt"

func dispacher(request Request) (Response, error) {

	switch request.Method {
	case "redis":
		var response = NewResponse()

		data, err := RedisAction(request.Action, request.Args)
		if err != nil {
			fmt.Println(err)
		}
		response.Data = data
		return response, nil
	}

	return Response{}, nil
}
