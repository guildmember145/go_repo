package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
		for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tiempoTranscurrido := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoTranscurrido)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "Server Down =C")
	}else {
		fmt.Println(servidor, "Server UP!")
	}

}