package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/dalet-oss/terraform-provider-iris/iris"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

var version = "was not built correctly" // set via the Makefile

func main() {
	versionFlag := flag.Bool("version", false, "print version information and exit")
	flag.Parse()
	if *versionFlag {
		printVersion(os.Stdout)
		os.Exit(0)
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: iris.Provider,
	})
}

func printVersion(writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "%s %s\n", os.Args[0], version)
	return err
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
