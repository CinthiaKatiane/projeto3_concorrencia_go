package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
)

type jogador struct{

	nome string

}

type partida struct{

	j1 jogador
	j2 jogador
	game int
	set int
	match int
}

func jogar(c chan string, p jogador) {

	rand.Seed(time.Now().UnixNano())
	shot := rand.Intn(10)

	for shot > -1 {

		switch {
		case shot < 3:
			c <- p.nome+ ": errou"
		case shot >= 3:
			c <- p.nome + ": acertou"
		}

		rand.Seed(time.Now().UnixNano())
		shot = rand.Intn(10)
	}
}

func score(c chan string, p partida) {

	sc_i := 0
	sc_o := 0

	for sc_i < p.game && sc_o < p.game{

		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

		if msg == p.j1.nome + ": errou"{
			sc_o +=1
			fmt.Printf("%v %v x %v %v \n", p.j1.nome, sc_i, p.j2.nome, sc_o)
		} else if msg == p.j2.nome + ": errou" {
			sc_i +=1
			fmt.Printf("%v %v x %v %v \n", p.j1.nome, sc_i, p.j2.nome, sc_o)
		} else {
			continue
		}
	}

	fmt.Printf("\nResultado final:\n")
	fmt.Printf("%v %v x %v %v \n", p.j1.nome, sc_i, p.j2.nome, sc_o)
	os.Exit(0)

}

func main() {

	var c = make(chan string)
	var points int

	fmt.Printf("Defina os números de pontos:")
	fmt.Scanln(&points)

	jogador1 := jogador{"Nadal"}
	jogador2 := jogador{"Guga"}

	partida := partida{jogador1, jogador2,points,1, 1}

	go jogar(c, jogador1)
	go jogar(c, jogador2)
	go score(c, partida)

	var input string
	fmt.Scanln(&input)

}