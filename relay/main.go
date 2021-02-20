package relay

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	socks5 "github.com/armon/go-socks5"
	yamux "github.com/hashicorp/yamux"
)

func New(tunnelAddr string, maxFailedConnections int) {
	config := yamux.DefaultConfig()

	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	failedTunnelConnections := 0

	if err != nil {
		log.Println(err)
		return
	}
	for {
		log.Println("[INFO] wait for tunnel")
		// Connect TLS socket wrapped with yamux
		conn, err := connectTunnel(tunnelAddr)
		if err != nil {
			log.Printf("[ERR] Error connecting tunnel (%d): %s", failedTunnelConnections, err.Error())
			failedTunnelConnections = failedTunnelConnections + 1
			if failedTunnelConnections > maxFailedConnections {
				return
			}
			fc := time.Duration(failedTunnelConnections)
			time.Sleep(fc * fc * 100 * time.Millisecond)
			continue
		}
		failedTunnelConnections = 0
		session, err := yamux.Client(conn, config)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("[INFO] wait for stream")
		// Accept new streams and pass them to a goroutine
		for {
			stream, err := session.Accept()
			if err != nil {
				log.Println("[ERR] Error accept new stream: " + err.Error())
				break
			}

			if err != nil {
				log.Println("[ERR] Error setting up SOCKS server: " + err.Error())
				stream.Close()
				continue
			}

			log.Println("[INFO] Send Stream to SOCKS")
			go server.ServeConn(stream)
		}
	}
}

func connectTunnel(serverHost string) (conn net.Conn, err error) {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	if conn, err = tls.Dial("tcp", serverHost, conf); err != nil {
		log.Println(err)
		return
	}

	return
}
