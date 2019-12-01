package main

import (
	"fmt"
	"io/ioutil"

	"github.com/containers/image/v5/signature"
	"gopkg.in/alecthomas/kingpin.v2"
)

func verifyCmd(app *kingpin.Application) {
	opts := &verifyOptions{}
	command := app.Command("verify", "Verify a Docker image.").Action(opts.run)

	command.Arg("manifest", "Path to a file containing the image manifest.").
		Required().
		StringVar(&opts.manifestPath)

	command.Arg("docker-reference", "Docker reference to identify the image with.").
		Required().
		StringVar(&opts.dockerReference)

	command.Arg("key-fingerprint", "GPG key identity to use for signing.").
		Required().
		StringVar(&opts.fingerprint)

	command.Arg("signature", "Path to the signature.").
		Required().
		StringVar(&opts.signaturePath)
}

type verifyOptions struct {
	manifestPath    string
	dockerReference string
	fingerprint     string
	outputPath      string
	signaturePath   string
}

func (opts *verifyOptions) run(ctx *kingpin.ParseContext) error {
	unverifiedManifest, err := ioutil.ReadFile(opts.manifestPath)
	if err != nil {
		return fmt.Errorf("Error reading manifest from %s: %v", opts.manifestPath, err)
	}
	unverifiedSignature, err := ioutil.ReadFile(opts.signaturePath)
	if err != nil {
		return fmt.Errorf("Error reading signature from %s: %v", opts.signaturePath, err)
	}

	mech, err := signature.NewGPGSigningMechanism()
	if err != nil {
		return fmt.Errorf("Error initializing GPG: %v", err)
	}
	defer mech.Close()
	sig, err := signature.VerifyDockerManifestSignature(unverifiedSignature, unverifiedManifest, opts.dockerReference, mech, opts.fingerprint)
	if err != nil {
		return fmt.Errorf("Error verifying signature: %v", err)
	}

	fmt.Printf("Signature verified, digest %s\n", sig.DockerManifestDigest)
	return nil
}
