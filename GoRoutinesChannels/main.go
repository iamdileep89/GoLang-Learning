package main

import (
	"fmt"
	"net/http"
	"time"
)

// This is a blocking channelling, since the channel is wating to hear from the go routine
// func main() {
// 	c := make(chan string)
// 	c <- "Hi there!"
// }

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.in",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// New Loop syntax for channel
	for l := range c {

		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
