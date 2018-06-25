package main

import (
	"math/rand"
	"time"
	"fmt"
	"os"
)

/* Foram definidos dois structs, o de jogador apresenta o nome do jogador
e os seus pontos marcados; o de partida apresenta as informações da partida
que esta sendo jogada, tais informações são configuraveis no inicio da partida*/

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

/* a funcao jogar eh executada pelos jogadores e informa se o jogador acertou
ou errou a jogada, ela tem base em um valor aleatorio e informa atravez do canal o
resultado da jogada */

func jogar(c chan string, p jogador) {

	rand.Seed(time.Now().UnixNano())
	shot := rand.Intn(10)

	for {
		/*a chance de erro eh menor que a de acerto para que o jogo fique mais
		dinamico, o resultado do acerto ou erro eh enviado no canal*/
		switch {
		case shot < 3:
			c <- "\t"+p.nome+ ": errou"
		case shot >= 3:
			c <- "\t"+p.nome + ": acertou"
		}

		rand.Seed(time.Now().UnixNano())
		shot = rand.Intn(10)
	}
}

/*a funcao score calcula o placar para a opcao personalizada da partida
ela apenas le as informacoes passadas pelo canal e atribui os pontos, os games e os sets*/

func score(c chan string, p partida) {
	i := 0

	for !(p.j1.w_set > p.set && p.j1.w_set > p.j2.w_set+1) && !(p.j2.w_set > p.set && p.j2.w_set > p.j1.w_set+1) {

		p.j1.w_game = 0
		p.j2.w_game = 0
		i+=1
		fmt.Printf("Set %v\n", i)
		j := 0
		for !(p.j1.w_game > p.game && p.j1.w_game>p.j2.w_game+1) && !(p.j2.w_game > p.game && p.j2.w_game> p.j1.w_game+1){

			j+=1
			p.j1.w_pontos = 0
			p.j2.w_pontos = 0
			fmt.Printf("Game %v\n", j)
			for p.j1.w_pontos < p.pontos && p.j2.w_pontos < p.pontos {

				msg := <-c
				fmt.Println(msg)
				time.Sleep(time.Second * 1)

				if msg == "\t"+p.j1.nome+": errou" {
					p.j2.w_pontos += 1
					fmt.Printf("Game Atual: %v %v x %v %v \n", p.j1.nome, p.j1.w_pontos, p.j2.nome, p.j2.w_pontos)
				} else if msg == "\t"+p.j2.nome+": errou" {
					p.j1.w_pontos += 1
					fmt.Printf("Game Atual: %v %v x %v %v \n", p.j1.nome, p.j1.w_pontos, p.j2.nome, p.j2.w_pontos)
				} else {
					continue
				}
			}
			if p.j1.w_pontos > p.j2.w_pontos{
				p.j1.w_game +=1
			}else {
				p.j2.w_game +=1
			}
			fmt.Printf("Games: %v %v x %v %v \n", p.j1.nome, p.j1.w_game, p.j2.nome, p.j2.w_game)

		}
		if p.j1.w_game > p.j2.w_game {
			p.j1.w_set +=1
		}else {
			p.j2.w_set +=1
		}
		fmt.Printf("Set: %v %v x %v %v \n", p.j1.nome, p.j1.w_set, p.j2.nome, p.j2.w_set)

	}
	fmt.Printf("\nResultado final do match:\n")
	if (p.j1.w_set > p.j2.w_set){
		fmt.Printf("Vitória de %v \n", p.j1.nome)
	} else if (p.j2.w_set > p.j1.w_set) {
		fmt.Printf("Vitória de %v \n", p.j2.nome)
	} else{
		fmt.Printf("Empate!")
	}

	os.Exit(0)
}


/*a funcao score calcula o placar de forma simplificada, pois leva em conta apenas o pontos do game
ela apenas le as informacoes passadas pelo canal e atribui os pontos*/
func simple_score(c chan string, p partida) {
	for p.j1.w_pontos < p.pontos && p.j2.w_pontos < p.pontos {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

		if msg == "\t"+p.j1.nome+": errou" {
			p.j2.w_pontos += 1
			fmt.Printf("Game: %v %v x %v %v \n", p.j1.nome, p.j1.w_pontos, p.j2.nome, p.j2.w_pontos)
		} else if msg == "\t"+p.j2.nome+": errou" {
			p.j1.w_pontos += 1
			fmt.Printf("Game: %v %v x %v %v \n", p.j1.nome, p.j1.w_pontos, p.j2.nome, p.j2.w_pontos)
		} else {
			continue
		}
	}
	fmt.Printf("\nResultado final:\n")
	if (p.j1.w_pontos > p.j2.w_pontos){
		fmt.Printf("Vitória de %v \n", p.j1.nome)
	} else if (p.j2.w_pontos > p.j1.w_pontos) {
		fmt.Printf("Vitória de %v \n", p.j2.nome)
	} else{
		fmt.Printf("Empate!")
	}

	os.Exit(0)

}

func main() {

	var c = make(chan string)

	fmt.Printf("Rafael Nadal x Gustavo Kuerten\n")
	jogador1 := jogador{"Nadal",0,0,0,}
	jogador2 := jogador{"Guga",0,0,0}

	var op int
	fmt.Printf("Digite 1 se deseja executar apenas 1 game de 5 pontos \n" +
		"Digite 2 se deseja executar apenas 1 game e voce define o numero de pontos\n" +
		"Digite 3 se deseja personalizar completamente a partida\n")

	fmt.Scanln(&op)

	if op == 1{
		partida := partida{jogador1, jogador2,5,1, 1}

		go jogar(c, jogador1)
		go jogar(c, jogador2)
		go simple_score(c, partida)

		var input string
		fmt.Scanln(&input)
	} else if op == 2 {

		var points int
		fmt.Printf("Defina os números de pontos:")
		fmt.Scanln(&points)

		partida := partida{jogador1, jogador2,points,1, 1}

		go jogar(c, jogador1)
		go jogar(c, jogador2)
		go simple_score(c, partida)

		var input string
		fmt.Scanln(&input)

	} else if op == 3{

		var points int
		var games int
		var sets int

		fmt.Printf("Defina os números de pontos:")
		fmt.Scanln(&points)
		fmt.Printf("Defina os números de games:")
		fmt.Scanln(&games)
		fmt.Printf("Defina os números de sets:")
		fmt.Scanln(&sets)


		partida := partida{jogador1, jogador2,points,games, sets}

		go jogar(c, jogador1)
		go jogar(c, jogador2)
		go score(c, partida)

		var input string
		fmt.Scanln(&input)

	} else {
		fmt.Printf("Entrada inválida")
	}
}