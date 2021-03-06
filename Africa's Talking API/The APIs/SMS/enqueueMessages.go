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

	// Define the recipient numbers in a comma separated string
	// Numbers should be in international format as shown
	var to string = "+254711XXXYYY,+254733YYYZZZ"

	// And of course we want our recipients to know what we really do
	var message string = "I'm a lumberjack and its ok, I sleep all night and I work all day"

	// enqueue flag is used to queue messages incase you are sending a high volume.
    // The default value is 0.
	var enqueue int = 1;

	// Building the post
	at := url.Values{}
	at.Set("username", username)
	at.Set("to", to)
	at.Set("message", message)
	at.Set("enqueue", enqueue)

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