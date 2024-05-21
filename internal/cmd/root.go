package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var debug bool

func SetVersionInfo(versionStr, commitStr string) {
	rootCmd.Version = fmt.Sprintf("%s-%s", versionStr, commitStr)
}

var rootCmd = &cobra.Command{
	Use:   "groxi",
	Short: "An all inclusive gorocks implementation",
	Long:  `Just drop the binary and run it. Certificates are auto generated.`,
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug mode")

}
