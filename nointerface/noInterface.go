package main

import "fmt"
//codigo sin utilizar interfaces, se repite mucho para realizar algo.
type perro struct{}

type pez struct{}

type pajaro struct{}

func (perro) caminar() string {
	return "Soy un perro y camino"
}

func (pez) nadar() string {
	return "Soy un pez y puedo nadar"
}

func (pajaro) volar() string {
	return "Soy un pajaro y puedo volar"
}

func moverPerro(dog perro) {
	fmt.Println(dog.caminar())
}

func nadarPez(fish pez){
	fmt.Println(fish.nadar())
}

func volarPajaro(bird pajaro)  {
	fmt.Println(bird.volar())
}

func main() {

	dog := perro{}
	moverPerro(dog)
	fish := pez{}
	nadarPez(fish)
	bird := pajaro{}
	volarPajaro(bird)
}