package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("clerk", "Command line utility that can create and verify signatures of a Docker image using GPG.")
	app.Version("0.0.1")

	signCmd(app)
	verifyCmd(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
