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
	voiceURL string = "https://voice.africastalking.com:443/call"
)

var username string = "kanyi"
var apikey string = "9eb01a98435e1d0cff3e8f6e5c408fa330add662bfcc0a7e9ddcd09b2a0c62fa"

func main() {
	call()
}

func call(){

	// Define the recipient numbers in a comma separated string
	// Numbers should be in international format as shown

	var from string = "+254711082524"

	// And of course we want our recipients to know what we really do
	var to string = "+254790807760"

	// Building the post
	at := url.Values{}
		at.Set("username", username)
		at.Set("from", from)
		at.Set("to", to)

	// Values.Encode() encodes the post values into "URL encoded"
	postparams := at.Encode()

	// Perform the post
	req, err := http.NewRequest("POST", voiceURL, strings.NewReader(postparams))
	
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

	fmt.Println(string(data))
	
}