package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

type JsonResponse struct {
	Id   uuid.UUID              `json:"id"`
	Err  string                 `json:"err"`
	Data map[string]interface{} `json:"data"`
}

func (r *JsonResponse) Set(value interface{}) {
	typeName := reflect.TypeOf(value).Name()
	r.Data[typeName] = value
}

func (r *JsonResponse) Get(value interface{}) error {
	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		return errors.New("value is not pointer")
	}

	typeNamePackage := reflect.TypeOf(value).String()
	typeName := strings.Split(typeNamePackage, ".")[1]
	valueInterface, ok := r.Data[typeName]

	if !ok {
		return errors.New(fmt.Sprintf("%s not in json-rpc", typeName))
	}

	valueByte, err := json.Marshal(valueInterface)
	if err != nil {
		return err
	}

	return json.Unmarshal(valueByte, value)
}
