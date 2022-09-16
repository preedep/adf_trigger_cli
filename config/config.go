package config

import (
	"encoding/json"
	"os"
)

type DataType string

/*
 */
const (
	STRING DataType = "string"
	LIST   DataType = "list"
)

type Parameter struct {
	Key   string   `json:"key"`
	Value string   `json:"value"`
	Type  DataType `json:"type"`
}
type Parameters struct {
	Params []Parameter `json:"params"`
}

func ReadParametersFile(file string) (*Parameters, error) {
	readFile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var p Parameters
	err = json.Unmarshal(readFile, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
