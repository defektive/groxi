package main

import (
	"flag"
	"fmt"
	"github.com/defektive/groxi/common"
	"github.com/defektive/groxi/server"
)

var tunnelAddr = flag.String("t", "127.0.0.1:8081", "Address to accept relay connections on.")
var socksAddr = flag.String("s", "127.0.0.1:1080", "Address to accept socks connections on.")
var version = flag.Bool("v", false, "prints groxi version")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("groxi v%s\n", common.Version)
		return
	}

	server.New(*tunnelAddr, *socksAddr)
}
