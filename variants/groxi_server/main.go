package main

import (
	"flag"
	"github.com/defektive/groxi/server"
)

var tunnelAddr = flag.String("tunnel", "127.0.0.1:8081", "The bind address on which to accept tunnel connections")
var socksAddr = flag.String("socks", "127.0.0.1:1080", "The bind address on which to listen for socks clients")

func main() {
	flag.Parse()
	server.New(*tunnelAddr, *socksAddr)
}
