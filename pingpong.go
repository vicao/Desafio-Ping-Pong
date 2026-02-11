package main

//Criar um progroma em Go que trabalhe com dois pacotes e concorrência.
//Escreva um código em Go utilizando todo o conhecimento adquirido até o momento!
// E nesse código você precisará, baseado em nossas aulas de concorrência,
//utilizar canais e goroutines para que o seu programa exiba, em alternância, as palavras "ping" e "pong".
import (
	"fmt"
	"time"
)

func main() {
	pingCh := make(chan struct{})
	pongCh := make(chan struct{})
	msg := make(chan string)

	go Ping(pingCh, pongCh, msg)
	go Pong(pongCh, pingCh, msg)
	go Imprimir(msg)

	// inicia com ping
	pingCh <- struct{}{}

	var entrada string
	fmt.Scanln(&entrada)
}

func Ping(pingCh, pongCh chan struct{}, msg chan string) {
	for {
		<-pingCh
		msg <- "ping"
		pongCh <- struct{}{}
	}
}

func Pong(pongCh, pingCh chan struct{}, msg chan string) {
	for {
		<-pongCh
		msg <- "pong"
		pingCh <- struct{}{}
	}
}

func Imprimir(msg chan string) {
	for {
		m := <-msg
		fmt.Println(m)
		time.Sleep(time.Second * 1)
	}
}
