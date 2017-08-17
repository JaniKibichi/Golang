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
	sendMessageURL string = "https://api.africastalking.com/version1/messaging"
)

var username string = ""
var apikey string = ""

func main() {
	sendMessage()
}

func sendMessage(){

	// Define the recipient numbers
	var to string = "+254711XXXYYY"

	// Specify your premium shortcode and keyword
	var shortcode string = "XXXXX"
	var keyword string = "premiumKeyword"

	// Set the bulkSMSMode flag to 0 so that the subscriber gets charged
    var bulkSMSMode int = 0;

	// retryDurationInHours: The numbers of hours our API should retry to send the message 
    // before giving a failure status incase it doesn't go through. It is optional
     // incase it doesn't go through. It is optional
        
    var retryDurationInHours string = "No. of hours to retry";
        
    var message string = "Get your daily message and thats how we roll.";

	// Building the post
	at := url.Values{}
		at.Set("username", username)
		at.Set("to", to)
		at.Set("message", message)
		at.Set("shortcode", shortcode)
		at.Set("keyword", keyword)
		at.Set("bulkSMSMode", bulkSMSMode)
		at.Set("retryDurationInHours", retryDurationInHours)

	// Values.Encode() encodes the post values into "URL encoded"
	postparams := at.Encode()

	// Perform the post
	req, err := http.NewRequest("POST", sendMessageURL, strings.NewReader(postparams))
	
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
	fmt.Printf(string(data))

}