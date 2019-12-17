package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkURL(url, c)
	}

	for range urls {
		fmt.Println(<-c)
	}

}

func checkURL(url string, c chan string) {
	_, error := http.Get(url)

	if nil != error {
		c <- url + ".... is down!"
		return
	}
	c <- url + ".... is up!"
}
