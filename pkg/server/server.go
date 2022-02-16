package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dkharms/json-rpc/pkg/procedure"
	"log"
	"net/http"
	"reflect"
)

type JsonRequest struct {
	Data map[string]interface{} `json:"data"`
}

func (r *JsonRequest) Get(value interface{}) error {
	typeName := reflect.TypeOf(value).Name()
	valueInterface, ok := r.Data[typeName]
	if !ok {
		return errors.New(fmt.Sprintf("%s not in json", typeName))
	}

	valueByte, err := json.Marshal(valueInterface)
	if err != nil {
		return err
	}

	return json.Unmarshal(valueByte, value)
}

type JsonResponse struct {
	Err  error                  `json:"err"`
	Data map[string]interface{} `json:"data"`
}

func (r *JsonResponse) Set(value interface{}) {
	typeName := reflect.TypeOf(value).Name()
	r.Data[typeName] = value
}

type ProcedureRecord [2]string

type Server struct {
	l *log.Logger
	p map[ProcedureRecord]procedure.Procedure
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement method for procedure routing
}

func New(l *log.Logger) *Server {
	return &Server{l: l}
}

func (s *Server) AddProcedure(userProcedure procedure.Procedure) {
	procedureRecord := ProcedureRecord{userProcedure.Name, userProcedure.Version}
	s.p[procedureRecord] = userProcedure
}

func (s *Server) Run(port string) error {
	return http.ListenAndServe(port, s)
}
