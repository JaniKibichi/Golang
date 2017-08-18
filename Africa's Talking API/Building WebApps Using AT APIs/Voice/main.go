package main

import (
    "fmt"
    "html/template"
	"io/ioutil"
    "log"
	"net/url"
    "net/http"
    "strings"
)

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("makecall.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("phoneN:", r.Form["phoneN"])

		var formdata []string = r.Form["phoneN"]

        for _,phonenumbers := range formdata{
			call(phonenumbers)
		}
		
		fmt.Println("Calling...")
    }
}

func call(phonenumber string) {

	const voiceURL string = "https://voice.africastalking.com:443/call"
	const username string = ""
	const apikey string = ""

	// Your Africa's Talking Virtual Number
	var from string = ""
	// Building the post
	at := url.Values{}
		at.Set("username", username)
		at.Set("from", from)
		at.Set("to", phonenumber)

	// Values.Encode() encodes the post values into "URL encoded"
	postparams := at.Encode()

	// Perform the post
	req, err := http.NewRequest("POST", voiceURL, strings.NewReader(postparams))
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey",apikey)
	
	c := &http.Client{}
	resp, err := c.Do(req)
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		
	}
	
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		
	}

	fmt.Println(string(data))

	
}

func main() {
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil) // Setting listening port accessible on http://localhost:9090/login
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}