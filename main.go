package main

import (
	"os"
)

func main() {
	c := &CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(c.Run(os.Args))
}
