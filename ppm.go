package main

import (
	"fmt"
	"github.com/ppm-org/ppm/cmd"
	"os"
	"path/filepath"
)

const version = "0.0.1"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		_, name := filepath.Split(os.Args[0])
		fmt.Printf("%s v%s\n", name, version)
		fmt.Println()
		printHelp()
		return
	}
	routeCommand(args)
}
func printHelp() {
	_, name := filepath.Split(os.Args[0])
	fmt.Printf("Usage: %s <command>\n", name)
	fmt.Println()
	fmt.Println("where <command> is one of:")
	fmt.Print("    ")
	fmt.Print("init (i)")
	fmt.Print(", add (a)")
	fmt.Print(", download (dl)")
	fmt.Println()
}
func routeCommand(args []string) {
	switch args[0] {
	case "init":
	case "i":
		cmd.Init(args[1:])
		break
	case "add":
	case "a":
		cmd.Add(args[1:])
		break
	case "download":
	case "dl":
		cmd.Download(args[1:])
		break
	}
}
