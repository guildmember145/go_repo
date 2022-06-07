package main

import (
	"fmt"
	"net/http"
	"time"
)


func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + "\nNo se encuentra disponible"
	}else {
		canal <- servidor + "\nSe encuentra disponible"
	}

}

// separacion

func main() {

	canal := make(chan string)
	inicio := time.Now()
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i:=0
	
	for i<2 {
		for _, servidor := range servidores {
		go revisarServidor(servidor,canal)
		}
		time.Sleep(1*time.Second)
		fmt.Println(<-canal)
		i++
	}

	tiempoTranscurrido := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoTranscurrido)
}

