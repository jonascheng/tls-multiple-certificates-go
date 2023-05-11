// a golang web server to serve echo requests with multiple certificates
package main

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// define a call back function GetCertificate
// this function will be called by the server when a new connection is established
// the server will pass the client hello message to this function
// the function should return a certificate based on the client hello message
func GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	// log the client hello message
	log.Printf("Client hello: %+v", clientHello)

	// specify rand seed with current time
	rand.Seed(time.Now().UnixNano())
	// generate a random number between 0 and 100
	randNum := rand.Intn(100)

	if randNum%2 == 0 {
		log.Printf("Using certificate v1")
		certPair, _ := tls.LoadX509KeyPair("server-v1.crt", "server-v1.key")
		return &certPair, nil
	} else {
		log.Printf("Using certificate v2")
		certPair, _ := tls.LoadX509KeyPair("server-v2.crt", "server-v2.key")
		return &certPair, nil
	}
}

func main() {
	// setup TLS configuration
	tlsConfig := &tls.Config{
		SessionTicketsDisabled:   true,
		PreferServerCipherSuites: true,
		// set minimum TLS version
		MinVersion: tls.VersionTLS12,
	}

	// load certificates from files
	// certPair_v1, _ := tls.LoadX509KeyPair("server-v1.crt", "server-v1.key")
	// certPair_v2, _ := tls.LoadX509KeyPair("server-v2.crt", "server-v2.key")
	// tlsConfig.Certificates = []tls.Certificate{certPair_v1, certPair_v2}

	// BuildNameToCertificate parses c.Certificates and builds c.NameToCertificate from the CommonName and SubjectAlternateName fields of each of the leaf certificates.
	// Deprecated: NameToCertificate only allows associating a single certificate with a given name. Leave that field nil to let the library select the first compatible chain from Certificates.
	// tlsConfig.BuildNameToCertificate()

	// set GetCertificate callback function
	tlsConfig.GetCertificate = GetCertificate

	// create a new http server
	server := &http.Server{
		Addr: "0.0.0.0:8443",
		// configure TLS
		TLSConfig: tlsConfig,
	}
	// register handler
	http.HandleFunc("/", handler)
	// start the server
	log.Fatal(server.ListenAndServeTLS("", ""))
}

// handler to serve echo requests
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
