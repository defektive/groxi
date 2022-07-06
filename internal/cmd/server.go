package cmd

import (
	"github.com/defektive/groxi/pkg/server"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var tunnelAddress string
var socksAddress string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server listen for socks clients and relay clients",
	Long:  `Setup 2 listeners one for socks clients, one for relay clients.`,
	Run: func(cmd *cobra.Command, args []string) {

		if debug == false {
			file, err := os.OpenFile("groxi.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			log.SetOutput(file)
		} else {
			log.SetOutput(os.Stdout)
		}

		server.New(tunnelAddress, socksAddress)
	},
}

func init() {

	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&tunnelAddress, "tunnel", "t", "0.0.0.0:8081", "The bind address on which to accept tunnel connections")
	serverCmd.Flags().StringVarP(&socksAddress, "socks", "s", "127.0.0.1:1080", "The bind address on which to accept SOCKSv5 clients")
}
