package main

import (
	"os"
	"fmt"
)

var vetError = fmt.Sprintf("", 1, 2, 3)

func main() {
	c := &CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(c.Run(os.Args))
}

func Hoge(){
	
}

var hoge = 1