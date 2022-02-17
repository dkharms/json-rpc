package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type JsonRequest struct {
	Data map[string]interface{} `json:"data"`
}

func (r *JsonRequest) Get(value interface{}) error {
	typeNamePackage := reflect.TypeOf(value).String()
	typeName := strings.Split(typeNamePackage, ".")[1]
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
	Err  string                 `json:"err"`
	Data map[string]interface{} `json:"data"`
}

func (r *JsonResponse) Set(value interface{}) {
	typeName := reflect.TypeOf(value).Name()
	r.Data[typeName] = value
}
