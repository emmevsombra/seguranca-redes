package main

import (
	"fmt"
	"io"
	"log"
	"myapp/certificate"
	"myapp/connection"
	"os"
	"strconv"
	"time"
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
	fmt.Printf("time before:::  %d\n", timeBefore)

	//cria conexão
	conn, err := connection.ConnectToServer(hostname, port, certName, privKeyName)
	if err != nil {
		log.Fatalf("Erro ao conectar com servidor remoto %v:%v : %v", hostname, port, err)
	}

	timeAfter := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("time after:::  %d\n", timeAfter)

	time := timeAfter - timeBefore
	fmt.Printf("transmission time:::  %d\n", time)

	//envio da mensagem
	message := "teste GO 0"
	n, err := io.WriteString(conn, message)

	if err != nil {
		log.Fatalf("Erro ao enviar a mensagem. %s", err)
	}
	log.Printf("Client enviou %q (%d bytes)", message, n)

	for i := 1; i < 5; i++ {
		message = "teste GO " + strconv.Itoa(i)
		n, err = io.WriteString(conn, message)
		if err != nil {
			log.Fatalf("Erro ao enviar a mensagem. %s", err)
		}
		log.Printf("Client enviou %q (%d bytes)", message, n)
	}

}
