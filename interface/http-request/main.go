package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")

	if nil != err {
		fmt.Println("Errro fetching the google url:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
