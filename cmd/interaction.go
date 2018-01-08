package cmd

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func AskText(msg string) (text string, err error) {
	fmt.Print(msg + ": ")
	text, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	Len := len(text)
	if Len > 0 && text[Len-1] == '\n' {
		text = text[:Len-1]
	}
	return
}
