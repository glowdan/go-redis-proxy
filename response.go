package main

type Response struct {
	//{"em":"", "ec":"", "data":""}
	Em   string `json:"em"`
	Ec   int    `json:"ec"`
	Data string `json:"data"`
}

func NewResponse() Response {
	return Response{
		Em:   "success",
		Ec:   200,
		Data: "",
	}
}
