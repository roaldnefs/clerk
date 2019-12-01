package main

import (
	"fmt"
	"io/ioutil"

	"github.com/containers/image/signature"
	"gopkg.in/alecthomas/kingpin.v2"
)

func signCmd(app *kingpin.Application) {
	opts := &signOptions{}
	command := app.Command("sign", "Sign a Docker image.").Action(opts.run)

	command.Arg("manifest", "Path to a file containing the image manifest.").
		Required().
		StringVar(&opts.manifestPath)

	command.Arg("docker-reference", "Docker reference to identify the image with.").
		Required().
		StringVar(&opts.dockerReference)

	command.Arg("key-fingerprint", "GPG key identity to use for signing.").
		Required().
		StringVar(&opts.fingerprint)

	command.Flag("output", "Output file.").
		Required().
		StringVar(&opts.outputPath)
}

type signOptions struct {
	manifestPath    string
	dockerReference string
	fingerprint     string
	outputPath      string
}

func (opts *signOptions) run(ctx *kingpin.ParseContext) error {
	manifest, err := ioutil.ReadFile(opts.manifestPath)
	if err != nil {
		return fmt.Errorf("Error reading %s: %v\n", opts.manifestPath, err)
	}

	mech, err := signature.NewGPGSigningMechanism()
	if err != nil {
		return fmt.Errorf("Error initializing GPG: %v\n", err)
	}
	defer mech.Close()

	signature, err := signature.SignDockerManifest(manifest, opts.dockerReference, mech, opts.fingerprint)
	if err != nil {
		return fmt.Errorf("Error creating signature: %v\n", err)
	}

	if err := ioutil.WriteFile(opts.outputPath, signature, 0644); err != nil {
		return fmt.Errorf("Error writing signature to %s: %v", opts.outputPath, err)
	}
	return nil
}
