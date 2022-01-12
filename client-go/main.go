package main

import (
	"fmt"
	"io"
	"log"
	"time"
	"myapp/certificate"
	"myapp/connection"
	"os"
)

var CLIENT_NAME = "CLIENT_GO"

func main() {
	var (
		certName    = CLIENT_NAME + "_cert.pem"
		privKeyName = CLIENT_NAME + "_privkey.key"
		hostname    = "129.159.50.106"
		port        = 443
	)
	if _, err := os.Stat(certName); os.IsNotExist(err) {
		fmt.Println("Certificado não encontrado, gerando um novo certificado...")
		certificate.GenerateSelfSignedCert(CLIENT_NAME, 2048, certName, privKeyName)
	}

	timeBefore := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("time before:::  %d\n",timeBefore)

	conn, err := connection.ConnectToServer(hostname, port, certName, privKeyName)
	if err != nil {
		log.Fatalf("Erro ao conectar com servidor remoto %v:%v : %v", hostname, port, err)
	}

	timeAfter := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("time after:::  %d\n",timeAfter)

	time :=  timeAfter - timeBefore
	fmt.Printf("transmission time:::  %d\n",time)

	//envio da mensagem
	message := "teste GO"
	n, err := io.WriteString(conn, message)

	if err != nil {
		log.Fatalf("Erro ao enviar a mensagem. %s", err)
	}
	log.Printf("Client enviou %q (%d bytes)", message, n)


}
