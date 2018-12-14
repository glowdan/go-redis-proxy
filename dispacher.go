package main

import "fmt"

func dispacher(request Request) (Response, error) {

	switch request.Method {
	case "redis":
		var response = NewResponse()

		data, err := RedisAction(request.Action, request.Args)
		if err != nil {
			fmt.Println(err)
			response.Em = err.Error()
			response.Ec = 400
		}
		response.Data = data
		return response, nil
	}

	return Response{Ec: 405, Em: "Unsupported method ", Data: ""}, nil
}
