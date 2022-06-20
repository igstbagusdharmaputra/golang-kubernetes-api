package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// get data env

	endpoint := os.Getenv("APISERVER")
	token := os.Getenv("TOKEN")
	// set url endpoint
	url := endpoint + "/api/v1/namespaces/default/configmaps/golang-example-configmap"

	// set insecure url https
	// client := &http.Client{}
	// client.Transport = &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	// secure url https
	certFile, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Println(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certFile)

	tlsConfig := &tls.Config{RootCAs: caCertPool}
	tlsConfig.BuildNameToCertificate()

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	// request data
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	// add token on bearer
	request.Header.Add("Authorization", "Bearer"+" "+token)

	resp, err := client.Do(request)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	responsedata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var data map[string]interface{}
	json.Unmarshal(responsedata, &data)

	result := data["data"].(map[string]interface{})

	for key, value := range result {
		fmt.Println(key + "=" + value.(string))
	}
}
