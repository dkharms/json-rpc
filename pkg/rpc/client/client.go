package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dkharms/json-rpc/pkg/api"
	"github.com/google/uuid"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func Call(addr string, name api.ProcedureName, version api.ProcedureVersion, value interface{}) (*api.JsonResponse, error) {
	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		return nil, errors.New("value is not pointer")
	}

	typeNamePackage := reflect.TypeOf(value).String()
	typeName := strings.Split(typeNamePackage, ".")[1]
	request := api.JsonRequest{Id: uuid.New(), Data: map[string]interface{}{typeName: value}}

	jData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	post, err := http.Post(fmt.Sprintf("http://%s/%s?version=%s", addr, name, version), "application/json", bytes.NewReader(jData))
	if err != nil {
		return nil, err
	}

	responseData, _ := io.ReadAll(post.Body)
	jResponse := &api.JsonResponse{}
	err = json.Unmarshal(responseData, jResponse)
	if err != nil {
		return nil, err
	}

	return jResponse, nil
}
