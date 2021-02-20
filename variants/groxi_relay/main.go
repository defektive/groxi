package main

import (
	"flag"
	"github.com/defektive/groxi/relay"
)

var tunnelAddr = flag.String("tunnel", "127.0.0.1:8081", "The bind address on which to accept tunnel connections")
var maxFailedConnections = flag.Int("fail", 30, "The number of connections to try before giving up. 30 is about 15 minutes. NOTE: exponential falloff on retries. sleepMilis(failedAttemptCount * failedAttemptCount * 100)")

func main() {
	flag.Parse()
	relay.New(*tunnelAddr, *maxFailedConnections)
}
