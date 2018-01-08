package main

import (
	"fmt"
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
		cmdInit()
		break
	case "add":
	case "a":
		cmdAdd()
		break
	case "download":
	case "dl":
		cmdDownload()
		break
	}
}
func cmdInit() {
	fmt.Println("init not impl")
}
func cmdAdd() {
	fmt.Println("add not impl")
}
func cmdDownload() {
	fmt.Println("download not impl")
}
