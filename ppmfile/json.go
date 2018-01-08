package ppmfile

import (
	"encoding/json"
)

func ToJSON(ppmfile Package) ([]byte, error) {
	return json.Marshal(ppmfile)
}

func FromJSON(jsonBytes []byte) (res Package, err error) {
	err = json.Unmarshal(jsonBytes, &res)
	return
}

type Package struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Hash    string    `json:"hash"`
	Deps    []Package `json:"deps"`
}
