package cmd

import (
	"github.com/defektive/groxi/relay"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var tunnelAddr string
var maxFailedConnections int

var relayCmd = &cobra.Command{
	Use:   "relay",
	Short: "Connects to a server and allows traffic to pass through it",
	Long: `Connects to a server and allows traffic to pass through it. For example:

To fail after one connection attempt:
    groxi relay -f 1 -t 10.0.1.10:8081


NOTE: exponential falloff on retries. sleepMilis(failedAttemptCount * failedAttemptCount * 100)
bash snippet to help you decide what number to use.
    last=0;for i in {1..30}; do me=$(echo "$last+($i^2*100)"|bc); echo "$me/1000/60"|bc; last=$me  ;done
`,
	Run: func(cmd *cobra.Command, args []string) {
		if debug == false {
			log.SetOutput(ioutil.Discard)
		} else {
			log.SetOutput(os.Stdout)
		}

		relay.New(tunnelAddr, maxFailedConnections)
	},
}

func init() {
	rootCmd.AddCommand(relayCmd)
	relayCmd.Flags().StringVarP(&tunnelAddr, "tunnel", "t", "127.0.0.1:8081", "The bind address on which to accept tunnel connections")

	relayCmd.Flags().IntVarP(&maxFailedConnections, "fail", "f", 30, "The number of connections to try before giving up. 30 is about 15 minutes.")
}
