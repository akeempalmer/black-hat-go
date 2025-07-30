package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
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
