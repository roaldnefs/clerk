package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

func verifyCmd(app *kingpin.Application) {
	opts := &verifyOptions{}
	command := app.Command("verify", "Verify a Docker image.").Action(opts.run)

	command.Arg("signature", "Path to the signature.").
		Required().
		StringVar(&opts.signature)
}

type verifyOptions struct {
	signature string
}

func (opts *verifyOptions) run(ctx *kingpin.ParseContext) error {
	fmt.Printf("Would have run verify command.\n")
	return nil
}
