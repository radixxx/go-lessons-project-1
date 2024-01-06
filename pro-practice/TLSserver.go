package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

var PORT = ":1443"

type handler struct{}

func main() {
	caCert, err := os.ReadFile("client.crt")
	if err != nil {
		fmt.Println(err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	srv := &http.Server{
		Addr: PORT,
		//TODO : update + fix !
		Handler:   &handler{},
		TLSConfig: cfg,
	}
	fmt.Println("Listening to port number", PORT)
	fmt.Println(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func (h *handler) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world! \n"))
}
