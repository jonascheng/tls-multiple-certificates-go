// a golang web server to serve echo requests with multiple certificates
package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	// setup TLS configuration
	tlsConfig := &tls.Config{
		SessionTicketsDisabled:   true,
		PreferServerCipherSuites: true,
		// set minimum TLS version
		MinVersion: tls.VersionTLS12,
	}

	// load certificates from files
	certPair_v1, _ := tls.LoadX509KeyPair("server-v1.crt", "server-v1.key")
	certPair_v2, _ := tls.LoadX509KeyPair("server-v2.crt", "server-v2.key")
	tlsConfig.Certificates = []tls.Certificate{certPair_v1, certPair_v2}
	//tlsConfig.Certificates = []tls.Certificate{certPair_v2}

	// BuildNameToCertificate parses c.Certificates and builds c.NameToCertificate from the CommonName and SubjectAlternateName fields of each of the leaf certificates.
	// Deprecated: NameToCertificate only allows associating a single certificate with a given name. Leave that field nil to let the library select the first compatible chain from Certificates.
	// tlsConfig.BuildNameToCertificate()

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
