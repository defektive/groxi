package main

import (
	"flag"
	"fmt"

	"github.com/defektive/groxi/pkg/server"
)

var tunnelAddr = flag.String("t", "127.0.0.1:8081", "Address to accept relay connections on.")
var socksAddr = flag.String("s", "127.0.0.1:1080", "Address to accept socks connections on.")
var showVersion = flag.Bool("v", false, "Print groxi version")

var version = "v0.0.0"
var commit = "replace"

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("groxi v%s-%s\n", version, commit)
		return
	}

	server.New(*tunnelAddr, *socksAddr)
}
