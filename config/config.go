package config

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
type Configuration struct {
	Parameters []Parameter `json:"params"`
}
