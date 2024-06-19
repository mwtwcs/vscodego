package main


import (

	"fmt"

	"net/http"

	"io/ioutil"

)
func deferrequest() {
	url := "https://example.com"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sending request to", url)
	resp, err := client.Do(req)
	if err!= nil {
		fmt.Println(err)
		return
	}
	defer func() {
		fmt.Println("Closing connection...")
		resp.Body.Close()
	}()
	fmt.Println("Response status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response body:", string(body))

}