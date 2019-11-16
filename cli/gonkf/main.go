package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/facade"
)

func main() {
	facade.Execute(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
