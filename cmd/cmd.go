package main

import (
	"fmt"
	"github.com/dkharms/json-rpc/pkg/procedure"
	"github.com/dkharms/json-rpc/pkg/server"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func main() {
	s := server.New(nil)
	s.AddProcedure(procedure.New("GetSum", "@1",
		func(request *server.JsonRequest, response *server.JsonResponse) {
			sr := &SumRequest{}
			err := request.Get(sr)
			if err != nil {
				return
			}

			res := SumResponse{Result: sr.A + sr.B}
			response.Set(res)
			fmt.Println(request, response)
		}))
	s.Run(":8080")
}
