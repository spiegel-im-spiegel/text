package facade

import (
	"fmt"

	"github.com/spf13/cobra"
)

//newVersionCmd returns cobra.Command instance for version sub-command
func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of " + Name,
		Long:  "Print the version number of " + Name,
		RunE: func(cmd *cobra.Command, args []string) error {
			cui.OutputErrln(fmt.Sprintf("%s %s", Name, Version))
			return nil
		},
	}

	return versionCmd
}
