package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
)

type jogador struct{

	nome string
	w_pontos int
	w_game int
	w_set int

}

type partida struct{

	j1 jogador
	j2 jogador
	pontos int
	game int
	set int
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

	for p.j1.w_set < p.set && p.j2.w_set < p.set {

		for p.j1.w_game < p.game && p.j2.w_game < p.game{

			pt1 := 0
			pt2 := 0

			for pt1 < p.pontos && pt2 < p.pontos {

				msg := <-c
				fmt.Println(msg)
				time.Sleep(time.Second * 1)

				if msg == p.j1.nome+": errou" {
					pt2 += 1
					fmt.Printf("%v %v x %v %v \n", p.j1.nome, pt1, p.j2.nome, pt2)
				} else if msg == p.j2.nome+": errou" {
					pt1 += 1
					fmt.Printf("%v %v x %v %v \n", p.j1.nome, pt1, p.j2.nome, pt2)
				} else {
					continue
				}
			}
			fmt.Printf("\nResultado final do game:\n")
			fmt.Printf("%v %v x %v %v \n", p.j1.nome, pt1, p.j2.nome, pt2)

		}
	}
	os.Exit(0)
}



func main() {

	var c = make(chan string)
	var points int

	fmt.Printf("Defina os números de pontos:")
	fmt.Scanln(&points)

	jogador1 := jogador{"Nadal",0,0,0,}
	jogador2 := jogador{"Guga",0,0,0}

	partida := partida{jogador1, jogador2,points,2, 2}

	go jogar(c, jogador1)
	go jogar(c, jogador2)
	go score(c, partida)

	var input string
	fmt.Scanln(&input)

}