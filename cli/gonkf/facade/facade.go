package facade

import (
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli"
)

var (
	//Name is applicatin name
	Name = "gonkf"
	//Version is version for applicatin
	Version string
)

var (
	cui = gocli.NewUI() //CUI instance
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: Name,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no command")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ui *gocli.UI, args []string) (exit ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = ExitAbnormal
		}
	}()

	//execution
	cui = ui
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())
	exit = ExitNormal
	if err := rootCmd.Execute(); err != nil {
		exit = ExitAbnormal
	}
	return
}

func init() {
}
