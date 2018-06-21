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
	i := 0
	j := 0
	for p.j1.w_set < p.set && p.j2.w_set < p.set {
		i = 0
		j++
		p.j1.w_game = 0
		p.j2.w_game = 0

		for p.j1.w_game < p.game && p.j2.w_game < p.game{
			i++
			p.j1.w_pontos = 0
			p.j2.w_pontos = 0

			for p.j1.w_pontos < p.pontos && p.j2.w_pontos < p.pontos {

				msg := <-c
				fmt.Println(msg)
				time.Sleep(time.Second * 1)

				if msg == p.j1.nome+": errou" {
					p.j2.w_pontos += 1
					fmt.Printf("%v %v x %v %v \n", p.j1.nome, p.j1.w_pontos, p.j2.nome, p.j2.w_pontos)
				} else if msg == p.j2.nome+": errou" {
					p.j1.w_pontos += 1
					fmt.Printf("%v %v x %v %v \n", p.j1.nome, p.j1.w_pontos, p.j2.nome, p.j2.w_pontos)
				} else {
					continue
				}
			}

			if p.j1.w_pontos > p.j2.w_pontos {
				p.j1.w_game += 1
			} else {
				p.j2.w_game += 1
			}
			fmt.Printf("\nResultado final do game %v:\n", i)
			fmt.Printf("%v %v x %v %v \n\n", p.j1.nome, p.j1.w_pontos, p.j2.nome, p.j2.w_pontos)

		}

		if p.j1.w_game > p.j2.w_game {
			p.j1.w_set += 1
		} else {
			p.j2.w_set += 1
		}
		fmt.Printf("\nResultado final do set %v:\n", j)
		fmt.Printf("%v %v x %v %v \n\n", p.j1.nome, p.j1.w_game, p.j2.nome, p.j2.w_game)

	}
	fmt.Printf("\nResultado final do match:\n")
	fmt.Printf("%v %v x %v %v \n\n", p.j1.nome, p.j1.w_set, p.j2.nome, p.j2.w_set)
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