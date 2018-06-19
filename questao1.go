package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
)

func pinger(c chan string) {
	shot := (rand.Intn(10))
	for shot > -1 {

		switch {
		case (shot < 3):
			for i := 0; i<1; i++ {
				c <- "erro ping"
			}
		case (shot >= 3):
			for i := 0;i<1 ; i++ {
				c <- "ping"
			}
		}
		rand.Seed(time.Now().UnixNano())
		shot = (rand.Intn(10))
	}
}

func ponger(c chan string) {
	rand.Seed(time.Now().UnixNano())
	shot := (rand.Intn(10))
	for shot > -1 {
		switch {
		case (shot < 3):
			for i := 0;i<1 ; i++ {
				c <- "erro pong"
			}
		case (shot >= 3):
			for i := 0;i<1 ; i++ {
				c <- "pong"
			}
		}
		rand.Seed(time.Now().UnixNano())
		shot = (rand.Intn(10))
	}
}

func score(c chan string, p int) {
	sc_i := 0
	sc_o := 0
	for sc_i < p && sc_o < p{
		msg := <- c

		fmt.Println(msg)
		time.Sleep(time.Second * 1)

		if msg == "erro ping" {
			sc_o +=1
			fmt.Printf("Ping %v x Pong %v \n", sc_i, sc_o)
		} else if msg == "erro pong" {
			sc_i +=1
			fmt.Printf("Ping %v x Pong %v \n", sc_i, sc_o)
		} else {
			continue
		}
	}
	fmt.Printf("\nResultado final:\n")
	fmt.Printf("Ping %v x Pong %v \n", sc_i, sc_o)
	os.Exit(0)
	close(c)
}

func main() {
	var c chan string = make(chan string)

	var points int
	fmt.Scanln(&points)

	go pinger(c)
	go ponger(c)
	go score(c, points)

	var input string
	fmt.Scanln(&input)

}