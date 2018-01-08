package ppmfile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const Filename = "ppm.json"

var ErrorPPMFileNotFound = errors.New("ppm file not found")

func ToJSON(ppmfile Package) ([]byte, error) {
	return json.Marshal(ppmfile)
}

func FromJSON(jsonBytes []byte) (res Package, err error) {
	err = json.Unmarshal(jsonBytes, &res)
	return
}
func FindFile() (fullPath string, err error) {
	path := "."
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() == Filename {
			fullPath = path + Filename
			return
		}
	}
	err = ErrorPPMFileNotFound
	return
}

// check if the file is right here (not recursively)
func HasFile(path string) (exist bool, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, f := range files {
		if f.Name() == Filename {
			exist = true
			return
		}
	}
	exist = false
	return
}
func DefaultPackage() (pkg Package, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	_, pkg.Name = filepath.Split(dir)
	pkg.Version = "0.0.1"
	return
}

type Package struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Hash    string    `json:"hash"`
	Deps    []Package `json:"deps"`
}
