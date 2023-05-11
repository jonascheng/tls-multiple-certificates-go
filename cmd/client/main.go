// a golang web clinet to issue requests with trusted server certificate
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// setup TLS configuration
	tlsConfig := &tls.Config{
		// set minimum TLS version
		MinVersion: tls.VersionTLS12,
	}

	// load trusted CA certificates from file
	caCert, _ := ioutil.ReadFile("server-v1.crt")

	// create a certificate pool and add CA certificate
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// set trusted CA certificate pool
	tlsConfig.RootCAs = caCertPool

	// create a new http client
	client := &http.Client{
		// configure TLS
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// issue a request
	resp, err := client.Get("https://172.31.1.10:8443/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print response body
	log.Println(string(body))
}
