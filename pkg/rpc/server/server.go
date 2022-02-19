package server

import (
	"encoding/json"
	"fmt"
	"github.com/dkharms/json-rpc/pkg/api"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct {
	l *log.Logger
	p map[api.ProcedureName]*api.ProcedureMap
}

var defaultServer = server{
	l: log.New(os.Stdout, "server ", log.Ldate|log.Lshortfile),
	p: map[api.ProcedureName]*api.ProcedureMap{},
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	s.l.Printf("[%s] new connection from %s\n", request.Method, request.RemoteAddr)
	path := api.ProcedureName(request.URL.Path[1:])
	version := api.ProcedureVersion(request.URL.Query().Get("version"))
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

	jResponse.Id = jRequest.Id
	err = procedureImpl.Handler(jRequest, jResponse)
	if err != nil {
		jResponse.Err = fmt.Sprintf("error [%s] appeared while executing %s", err, path)
		return
	}
}

func New(l *log.Logger) *server {
	return &server{l: l, p: map[api.ProcedureName]*api.ProcedureMap{}}
}

func AddProcedure(userProcedure api.Procedure) {
	defaultServer.AddProcedure(userProcedure)
}

func (s *server) AddProcedure(userProcedure api.Procedure) {
	p, ok := s.p[userProcedure.Name]

	if !ok {
		p = &api.ProcedureMap{}
		s.p[userProcedure.Name] = p
	}

	p.Add(userProcedure)
}

func Run(addr string) error {
	return defaultServer.Run(addr)
}

func (s *server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}
