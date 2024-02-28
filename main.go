package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com"}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		msg := link + " is down!"
		fmt.Println(msg)
		c <- link

		return
	}
	msg := link + " is up!"
	fmt.Println(msg)
	c <- link

}
