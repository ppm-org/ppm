package cmd

import (
	"errors"
	"fmt"
	"github.com/ppm-org/ppm/ppmfile"
	"io/ioutil"
	"log"
)

var ErrorFileExist = errors.New("ppm file exist")

func Init(args []string) (err error) {
	exist, err := ppmfile.HasFile(".")
	if err != nil {
		log.Fatal(err)
		return
	}
	if exist {
		hasForceFlag := len(args) > 0 && isForceFlag(args[0])
		if !hasForceFlag {
			fmt.Println(ppmfile.Filename, "already exist! Please delete it if you want to reset the configuration.")
			err = ErrorFileExist
			return
		}
	}
	pkg, err := ppmfile.DefaultPackage()
	if err != nil {
		log.Fatalln(err)
		return
	}
	var text string

	text, err = AskText("package name (" + pkg.Name + ")")
	if err != nil {
		log.Fatalln(err)
		return
	}
	if text != "" {
		pkg.Name = text
	}

	text, err = AskText("version (" + pkg.Version + ")")
	if err != nil {
		log.Fatalln(err)
		return
	}
	if text != "" {
		pkg.Version = text
	}

	bs, err := ppmfile.ToJSON(pkg)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = ioutil.WriteFile(ppmfile.Filename, bs, 0644)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("saved to", ppmfile.Filename)

	return
}
func isForceFlag(s string) bool {
	return s == "--force" ||
		s == "-f"
}
