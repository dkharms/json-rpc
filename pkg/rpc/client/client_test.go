package client

import (
	"github.com/dkharms/json-rpc/pkg/api"
	"github.com/dkharms/json-rpc/pkg/rpc/server"
	"testing"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

func RunServer() {
	server.AddProcedure(api.Procedure{
		Name:    "GetSum",
		Version: "v1",
		Handler: func(request *api.JsonRequest, response *api.JsonResponse) error {
			sumRequest := SumRequest{}
			err := request.Get(&sumRequest)
			if err != nil {
				return err
			}
			response.Set(SumResponse{Sum: sumRequest.A + sumRequest.B})
			return nil
		},
	})
	server.Run(":8080")
}

func TestRequest(t *testing.T) {
	go RunServer()

	sumRequest := SumRequest{A: 30, B: 40}
	response, err := Call(":8080", "GetSum", "v1", &sumRequest)
	if err != nil {
		t.Fatalf("got error: %s", err)
	}
	sumResponse := SumResponse{}
	err = response.Get(&sumResponse)
	if err != nil {
		t.Fatalf("couldn't get value from json: %s", err)
	}
	if sumResponse.Sum != 70 {
		t.Fatalf("expected 70, got %d", sumResponse.Sum)
	}
}
