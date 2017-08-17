// Sending SMS on Africa's Talking API

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var username string = "MyAfricasTalkingUsername"
var apikey string = "MyAfricasTalkingApiKey"

func main() {
	fetchMessages(lastReceivedId)
}

func fetchMessages(lastReceivedId){
	const (
		sendMessageURL string = "https://api.africastalking.com/version1/messaging?username=",username,"&lastReceivedId=",lastReceivedId
	)

	// Perform the post
	req, err := http.NewRequest("GET", sendMessageURL, strings.NewReader())
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey",apikey)
	
	c := &http.Client{}
	resp, err := c.Do(req)
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Print the response from the Gateway
	fmt.Printf("%v\n", string(data))

}