package server

import (
	"encoding/json"
	"fmt"
	"github.com/dkharms/json-rpc/pkg/api"
	"github.com/dkharms/json-rpc/pkg/procedure"
	"io"
	"log"
	"net/http"
)

type Server struct {
	l *log.Logger
	p map[procedure.Name]*procedure.Map
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	s.l.Printf("[%s] new connection from %s\n", request.Method, request.RemoteAddr)
	path := procedure.Name(request.URL.Path[1:])
	version := procedure.Version(request.URL.Query().Get("version"))
	s.l.Printf("[%s] calling %s with version %s", request.RemoteAddr, path, version)

	jRequest, jResponse := &api.JsonRequest{}, &api.JsonResponse{Data: map[string]interface{}{}}
	defer func() {
		writer.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(jResponse)
		writer.Write(data)
	}()

	procedureMap, ok := s.p[path]
	if !ok {
		jResponse.Err = fmt.Sprintf("procedure %s wasn't found", path)
		return
	}

	procedureImpl, ok := procedureMap.Get(version)
	if !ok {
		jResponse.Err = fmt.Sprintf("procedure %s version %s wasn't found", path, version)
		return
	}

	readData, _ := io.ReadAll(request.Body)
	err := json.Unmarshal(readData, jRequest)
	if err != nil {
		jResponse.Err = fmt.Sprintf("can't parse json data %s", err)
		return
	}

	err = procedureImpl.Impl(jRequest, jResponse)
	if err != nil {
		jResponse.Err = fmt.Sprintf("error [%s] appeared while executing %s", err, path)
		return
	}
}

func New(l *log.Logger) *Server {
	return &Server{l: l, p: map[procedure.Name]*procedure.Map{}}
}

func (s *Server) AddProcedure(userProcedure procedure.Procedure) {
	_, ok := s.p[userProcedure.Name]

	if !ok {
		s.p[userProcedure.Name] = &procedure.Map{}
	}

	p := s.p[userProcedure.Name]
	p.Add(userProcedure)
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}
