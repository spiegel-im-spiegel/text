package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/facade"
)

func main() {
	os.Exit(facade.Execute(
		gocli.NewUI(
			gocli.Reader(os.Stdin),
			gocli.Writer(os.Stdout),
			gocli.ErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Int())
}
