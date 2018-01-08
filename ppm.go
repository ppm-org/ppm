package main

import (
	"fmt"
	"github.com/ppm-org/ppm/cmd"
	"os"
	"path/filepath"
)

const name = "ppm"
const version = "0.0.1"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printVersion()
		fmt.Println()
		printHelp()
		return
	}
	routeCommand(args)
}
func printVersion() {
	fmt.Printf("%s v%s\n", name, version)
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
		cmd.Init(args[1:])
	case "i":
		cmd.Init(args[1:])
	case "add":
		cmd.Add(args[1:])
	case "a":
		cmd.Add(args[1:])
	case "download":
		cmd.Download(args[1:])
	case "dl":
		cmd.Download(args[1:])
	default:
		printVersion()
		fmt.Fprint(os.Stderr, "command <", args[0], "> not found!\n")
		os.Exit(1)
		return
	}
}
