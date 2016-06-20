package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func EncodeBody(body interface{}) (*bytes.Buffer, error) {
	if body == nil {
		panic("Body argument should not be nil")
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buf, nil
}
