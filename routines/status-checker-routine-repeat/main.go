package main

import (
	"fmt"
	"net/http"
	"time"
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

	for url := range c {
		go func(url string) {
			time.Sleep(3 * time.Second)
			checkURL(url, c)
		}(url)
	}

}

func checkURL(url string, c chan string) {
	_, error := http.Get(url)
	if nil != error {
		fmt.Println(url, ".... is down!")
		c <- url
		return
	}
	fmt.Println(url, ".... is up!")
	c <- url
}
