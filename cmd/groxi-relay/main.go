package main

import (
	"flag"
	"fmt"

	"github.com/defektive/groxi/pkg/relay"
)

var maxFailedConnections = flag.Int("f", 137, "The number of connections to try before giving up. NOTE: Exponential fall off. 137 is about 24 hours. millis(numFailedConnections * numFailedConnections * 100)")
var tunnelAddr = flag.String("t", "127.0.0.1:8081", "Address to connect to server on.")
var showVersion = flag.Bool("v", false, "Print groxi version")

var version = "v0.0.0"
var commit = "replace"

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("groxi v%s-%s\n", version, commit)
		return
	}

	relay.New(*tunnelAddr, *maxFailedConnections)
}
