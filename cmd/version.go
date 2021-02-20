package cmd

import (
	"fmt"
	"github.com/defektive/groxi/common"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of groxi",
	Long:  `All software has versions. This is groxi's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("groxi v%s\n", common.Version)
	},
}
