package main

import (
	"fmt"
	 "net/http"
	 "time"
)

func main() {
	links:= []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	c:= make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l:= range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println(link, "is up!")
	} else {
		fmt.Println(link, "is down with status code:", resp.StatusCode)
	}
	c <- link
}