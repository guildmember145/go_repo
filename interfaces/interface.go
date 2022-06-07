package main

import "fmt"


type animal interface {
	mover() string
}



type perro struct{}

type pez struct{}

type pajaro struct{}

func (perro) mover() string {
	return "Soy un perro y camino"
}

func (pez) mover() string {
	return "Soy un pez y puedo nadar"
}

func (pajaro) mover() string {
	return "Soy un pajaro y puedo volar"
}

func moverAnimal(a animal){
	fmt.Println(a.mover())
}

func main() {

	dog := perro{}
	moverAnimal(dog)
	fish := pez{}
	moverAnimal(fish)
	bird := pajaro{}
	moverAnimal(bird)
}