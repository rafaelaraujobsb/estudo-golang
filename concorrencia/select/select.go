package main

import (
	"fmt"
	"time"

	"github.com/rafaelaraujobsb/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url2)

	select {
	case t1 := <-c1: // verifica se chegou uma informação no canal
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
	default:
		return "Sem resposta"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
	)

	fmt.Println(campeao)
}
