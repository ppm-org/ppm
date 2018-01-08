package cmd

import (
	"github.com/ppm-org/ppm/ppmfile"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	pkg := ppmfile.Package{
		Name:    "ppm",
		Version: "0.0.1",
		Deps:    []ppmfile.Package{},
	}
	t.Log("pkg:", pkg)
	bs, err := ppmfile.ToJSON(pkg)
	if err != nil {
		t.Error("Failed to convert ppmfile into JSON:", err)
		t.Fail()
		return
	}
	t.Log("bs:", string(bs))
	res, err := ppmfile.FromJSON(bs)
	if err != nil {
		t.Error("Failed to convert ppmfile from JSON:", err)
		t.Fail()
		return
	}
	t.Log("res:", res)
	if !reflect.DeepEqual(res, pkg) {
		t.Fail()
	}
}
