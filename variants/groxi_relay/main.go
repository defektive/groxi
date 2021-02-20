package main

import (
	"flag"
	"fmt"
	"github.com/defektive/groxi/common"
	"github.com/defektive/groxi/relay"
)

var maxFailedConnections = flag.Int("f", 30, "The number of connections to try before giving up. NOTE: Exponential fall off")
var tunnelAddr = flag.String("t", "127.0.0.1:8081", "Address to connect to server on.")
var version = flag.Bool("v", false, "Print groxi version")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("groxi v%s\n", common.Version)
		return
	}

	relay.New(*tunnelAddr, *maxFailedConnections)
}
