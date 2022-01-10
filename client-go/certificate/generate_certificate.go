package certificate

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
)

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

// Returns a PublicKey from a private key
func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// TODO: error return
func GenerateSelfSignedCert(clientName string, bitsOfKey int, certName, privKeyName string) {
	// Generate a private key for generating the certificate
	priv, err := rsa.GenerateKey(rand.Reader, bitsOfKey)

	if err != nil {
		log.Fatal(err)
	}

	// Template for a x509.Certificate
	// os.Hostname()
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:       []string{clientName}, // clientName
			CommonName:         "EMPRESA_CLIENTE",
			OrganizationalUnit: []string{hostname}, // ExtraInfos
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	out := &bytes.Buffer{}
	pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	ioutil.WriteFile(certName, out.Bytes(), 0644)

	out.Reset()

	pem.Encode(out, pemBlockForKey(priv))
	ioutil.WriteFile(privKeyName, out.Bytes(), 0644)
}

// Check if the certificate is valid
func CheckIfServerIsValid(certSent *x509.Certificate) error {
	const (
		server = "localhost"
	)

	verify := certSent.VerifyHostname(server)
	if verify == nil {
		return errors.New("certificate is not from server")
	} else {
		return nil
	}
}
