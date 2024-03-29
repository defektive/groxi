package main

import (
	"flag"
	"fmt"

	"github.com/defektive/groxi/internal"
	"github.com/defektive/groxi/pkg/relay"
)

var maxFailedConnections = flag.Int("f", 137, "The number of connections to try before giving up. NOTE: Exponential fall off. 137 is about 24 hours. millis(numFailedConnections * numFailedConnections * 100)")
var tunnelAddr = flag.String("t", "127.0.0.1:8081", "Address to connect to server on.")
var version = flag.Bool("v", false, "Print groxi version")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("groxi v%s\n", internal.Version)
		return
	}

	relay.New(*tunnelAddr, *maxFailedConnections)
}
