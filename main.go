package main

import (
	"fmt"
	"os"
	"os/user"
	"tungsten/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Headache++\n", user.Username)
	fmt.Printf("Send a command\n")
	repl.Start(os.Stdin, os.Stdout)
}
