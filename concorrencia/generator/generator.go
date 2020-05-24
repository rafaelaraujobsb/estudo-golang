package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GOOGLE I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			c <- string(html)
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.google.com", "https://www.youtube.com")
	t2 := titulo("https://www.cod3r.com.br", "https://www.amazon.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
