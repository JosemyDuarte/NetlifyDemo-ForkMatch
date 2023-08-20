package main

import (
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled code.
	Version = "dev"

	// id is the hostname of the machine.
	id, _ = os.Hostname()
)

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	app := NewApp(&config)

	if err = app.Run(); err != nil {
		panic(err)
	}
}
