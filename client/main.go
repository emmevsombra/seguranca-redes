package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var CLIENT_NAME = "CLIENT_1"

func main() {
	var (
		certName    = CLIENT_NAME + "_cert.pem"
		privKeyName = CLIENT_NAME + "_privkey.key"
		hostname    = "uece.br"
		port        = 443
	)
	if _, err := os.Stat(certName); os.IsNotExist(err) {
		fmt.Println("Certificado não encontrado, gerando um novo certificado...")
		certificate.GenerateSelfSignedCert(CLIENT_NAME, 2048, certName, privKeyName)
	}

	conn, err := connection.ConnectToServer(hostname, port, certName, privKeyName)
	if err != nil {
		log.Fatalf("Erro ao conectar com servidor remoto %v:%v : %v", hostname, port, err)
	}

	message := "teste"
	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("Erro ao enviar a mensagem. %s", err)
	}
	log.Printf("Client enviou %q (%d bytes)", message, n)
}
