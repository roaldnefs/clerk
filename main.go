package main

import(
  "fmt"
  "io/ioutil"

  "gopkg.in/alecthomas/kingpin.v2"
  "github.com/containers/image/v5/signature"
)

var (
  manifestPath    = kingpin.Arg("manifest", "Path to a file containing the image manifest.").Required().String()
  dockerReference = kingpin.Arg("docker-reference", "Docker reference to identify the image with.").Required().String()
  fingerprint     = kingpin.Arg("key-fingerprint", "GPG key identity to use for signing.").Required().String()
  outputPath      = kingpin.Flag("output", "Output file.").Required().String()
)

func main() {
  kingpin.Version("0.0.1")
  kingpin.Parse()

  manifest, err := ioutil.ReadFile(*manifestPath)
  if err != nil {
    fmt.Printf("Error reading %s: %v\n", *manifestPath, err)
    return
	}

  mech, err := signature.NewGPGSigningMechanism()
  if err != nil {
    fmt.Printf("Error initializing GPG: %v\n", err)
    return
  }
  defer mech.Close()

  signature, err := signature.SignDockerManifest(manifest, *dockerReference, mech, *fingerprint)
  if err != nil {
    fmt.Printf("Error creating signature: %v\n", err)
    return
  }

  if err := ioutil.WriteFile(*outputPath, signature, 0644); err != nil {
    fmt.Printf("Error writing signature to %s: %v", *outputPath, err)
    return
	}
}
