package main

import (
	"app"
	"os"
)

func main() {
	c := &app.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(c.Run(os.Args))
}
