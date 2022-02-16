package main

import (
	"github.com/dkharms/json-rpc/pkg/procedure"
	"github.com/dkharms/json-rpc/pkg/server"
	"log"
	"os"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func main() {
	l := log.New(os.Stdin, "server: ", log.Ldate|log.Lshortfile)
	s := server.New(l)
	s.AddProcedure(procedure.New("GetSum", "@1",
		func(request *server.JsonRequest, response *server.JsonResponse) error {
			sr := &SumRequest{}
			err := request.Get(sr)

			if err != nil {
				return err
			}

			res := SumResponse{Result: sr.A + sr.B}
			response.Set(res)

			return nil
		}))
	s.Run(":8080")
}
