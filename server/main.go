package server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"log"
	"math/big"
	"net"
	"sync"
	"time"

	"crypto/tls"
	"crypto/x509"
	yamux "github.com/hashicorp/yamux"
)

var tunnelAddress string
var socksAddress string

func New(tunnelAddress string, socksAddress string) {
	server, err := createTunnelServer(tunnelAddress)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	for {
		log.Println("Waiting for connections...")
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		handleConnection(conn, socksAddress)
	}
}

func createTunnelServer(host string) (server net.Listener, err error) {
	certPem, keyPem, err := certsetup()

	cer, err := tls.X509KeyPair(certPem.Bytes(), keyPem.Bytes())
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	server, err = tls.Listen("tcp", host, config)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func certsetup() (certPEM *bytes.Buffer, certPrivKeyPEM *bytes.Buffer, err error) {
	// set up our CA certificate
	// TODO: Randomize this :D
	log.Println("Creating Certificates...")

	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// create our private and public key
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	// create the CA
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, nil, err
	}

	// pem encode
	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})

	// set up our server certificate
	// TODO: Randomize this :D
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, nil, err
	}

	certPEM = new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	certPrivKeyPEM = new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})

	return
}

func handleConnection(conn net.Conn, socksHost string) {
	// Wrap connection in yamux
	session, err := yamux.Server(conn, nil)
	if err != nil {
		log.Println("[!] Error initializing yamux tunnel")
		return
	}
	defer conn.Close()
	defer session.Close()

	// Listen on SOCKSv5 server port
	socksServer, err := net.Listen("tcp", socksHost)
	if err != nil {
		log.Println("[!] Error listening for SOCKSv5 clients: " + err.Error())
		return
	}
	defer socksServer.Close()

	// Accept on socksv5 port; open new stream and start new goroutine to proxy
	log.Println("[+] Waiting for SOCKS clients on " + socksHost)
	for {
		client, err := socksServer.Accept()
		if err != nil {
			log.Println("[!] Error accepting connection from SOCKS client: " + err.Error())
			continue
		}
		stream, err := session.Open()
		if err != nil {
			log.Println("[+] Error opening new stream in tunnel: " + err.Error())
			return
		}
		go proxy(client, stream)
	}
	return
}

func proxy(c1 net.Conn, c2 net.Conn) {
	var wg sync.WaitGroup

	intProxy := func(a net.Conn, b net.Conn) {
		defer a.Close()
		defer b.Close()
		io.Copy(a, b)
		wg.Done()
	}

	go intProxy(c1, c2)
	wg.Add(1)

	go intProxy(c2, c1)
	wg.Add(1)

	wg.Wait()

	return
}
