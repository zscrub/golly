package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/zscrub/golly/repl"
)

func main() {
	user, err := user.Current()	
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %s | golly shell\n\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
