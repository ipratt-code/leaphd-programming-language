package main

import (
	"bufio"
	"fmt"
	"github.com/ipratt-code/hexal/toolchain"
	"os"
)

func main() {
	lic := make(chan string)
	lexer := toolchain.NewLexer(lic)

	go lexer.Run()

	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		t, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		lexer.In <- t
	}
}
