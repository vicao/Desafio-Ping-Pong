package main

//Criar um progroma em Go que trabalhe com dois pacotes e concorrência.
//Escreva um código em Go utilizando todo o conhecimento adquirido até o momento!
// E nesse código você precisará, baseado em nossas aulas de concorrência,
//utilizar canais e goroutines para que o seu programa exiba, em alternância, as palavras "ping" e "pong".
import "fmt"
import "time"

func Ping(c chan string) {//palavra reservada para criar um Canal : chan
	for i :=0;;i++{
		c <- "ping"//usado para enviar e receber mensagen para o Canal		
	}

}

func Pong(c chan string) {//palavra reservada para criar um Canal : chan
	for i :=0;;i++{
		c <- "pong"//usado para enviar e receber mensagen para o Canal		
	}

}

func Imprimir(c chan string) {
		for{
			msg := <-c
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
}

func main() {
	var c chan string = make(chan string) //cria um Canal do tipo string

	go Ping(c)
	go Imprimir(c)
	go Pong(c)
	go Imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)

}
