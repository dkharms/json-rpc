package main

import (
	"github.com/dkharms/json-rpc/pkg/api"
	"github.com/dkharms/json-rpc/pkg/rpc/server"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func RunServer() {
	server.AddProcedure(api.Procedure{
		Name:    "GetSum",
		Version: "@1",
		Handler: func(request *api.JsonRequest, response *api.JsonResponse) error {
			sr := &SumRequest{}

			err := request.Get(sr)
			if err != nil {
				return err
			}
			res := SumResponse{Result: sr.A + sr.B}

			response.Set(res)
			return nil
		}})
	server.Run(":8080")
}

func main() {
	RunServer()
}
