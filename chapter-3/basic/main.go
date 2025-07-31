package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	usingGetRequest()

}

func otherBasicRequests() {
	r1, err := http.Get("http://www.google.com/robots.txt")
	// Read response body. Notshown
	if err != nil {
		log.Fatal(err)
	}

	defer r1.Body.Close()

	r2, err := http.Head("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Red response body. Not shown.
	defer r2.Body.Close()

	form := url.Values{}

	form.Add("foo", "bar")

	// r3, err := http.Post("http://www.google.com/robots.txt", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))

	// Replace the above with PostForm below
	r3, err := http.PostForm("https://www.google.com/robots.txt", form)

	// Read response body. Not shown.

	defer r3.Body.Close()
}

func usingDeleteRequest() {
	// Seding a delete request.
	req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
	var client http.Client
	resp, err := client.Do(req)
	// Read response body and close.
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func usingPutRequest() {
	// Sending a Put Request
	form2 := url.Values{}
	form2.Add("foo", "bar")
	var client2 http.Client
	request, err := http.NewRequest(
		"PUT",
		"https://www.google.com/robots.txt",
		strings.NewReader(form2.Encode()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	resp2, err := client2.Do(request)
	// Read the response body and close.

	body2, err := io.ReadAll(resp2.Body)

	fmt.Println(string(body2))

}

func usingGetRequest() {
	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}

	// Print HTTP Status
	fmt.Println(resp.Status)

	// Read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(body))
	resp.Body.Close()
}
