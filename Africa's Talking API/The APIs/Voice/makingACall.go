// Sending SMS on Africa's Talking API

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
		makingCallURL string = "https://voice.africastalking.com/call"
	)

var username string = "MyAfricasTalkingUsername"
var apikey string = "MyAfricasTalkingApiKey"

func main() {
	fetchMessages(lastReceivedId)
}

func call(from string, to string){

	var from string = "myAfricasTalkingPhoneNumber"
	var to string = ""
	
	// Building the post
	at := url.Values{}
		at.Set("username", username)
		at.Set("from", from)
		at.Set("to", to)

	// Values.Encode() encodes the post values into "URL encoded"
	postparams := at.Encode()


	// Perform the request
	req, err := http.NewRequest("POST", makingCallURL, strings.NewReader(postparams))
	
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