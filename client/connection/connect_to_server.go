package connection

import (
	"crypto/tls"
	"errors"
	"fmt"
	"myapp/certificate"
)

// Returns a connection to the hostname:port
func ConnectToServer(hostname string, port int, certName, privKeyName string) (*tls.Conn, error) {
	const (
		protocol = "tcp"
	)
	// Load agent's self signed certificate
	cert, err := tls.LoadX509KeyPair(certName, privKeyName)
	if err != nil {
		return nil, errors.New("unable to load agent certificate " + certName)
	}

	// Configuration flags for TLS connection
	// InsecureSkipVerify need to be change to false in production realease
	config := tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cert}}
	config.MinVersion = tls.VersionTLS12
	config.MaxVersion = tls.VersionTLS13

	// TODO: tcp? why not http?
	// Connects to the server and
	conn, err := tls.Dial(protocol, hostname+":"+fmt.Sprint(port), &config)
	if err != nil {
		return nil, errors.New("unable to tls.Dial to " + hostname + ":" + fmt.Sprint(port))
	}
	//defer conn.Close() // Really needed?

	// Get server certificate and check if it belongs to Morphus
	state := conn.ConnectionState()
	certSent := state.PeerCertificates[0]
	err = certificate.CheckIfServerIsValid(certSent)
	if err != nil {
		return nil, errors.New("server certificate is not from Morphus")
	}

	return conn, nil
}
