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

	for _, url := range urls {
		go checkURL(url)
	}

}

func checkURL(url string) {
	_, error := http.Get(url)

	if nil != error {
		fmt.Println(url, ".... is down!")
		return
	}

	fmt.Println(url, ".... is up!")
}
