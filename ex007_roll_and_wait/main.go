package main

import (
	"fmt"
	"time"
	"math/rand"
)

type person struct {
	name string
	score int
}

func (player *person) play() {

	for {
		roll := rand.Intn(6) + 1
		wait := 6 - roll

		if (wait == 0) {
			fmt.Printf("%s (%d) rolled a %d, rolling again immediately\n", player.name, player.score, roll)
		} else {
			fmt.Printf("%s (%d) rolled a %d, waiting %d sec\n", player.name, player.score, roll, wait)
		}

		player.score = player.score + roll
		time.Sleep(time.Second * time.Duration(wait))
	}

}

func main() {
	timeUp := false

	players := []person{}
	players = append(players, person{name: "Ed", score: 0})
	players = append(players, person{name: "Natasha", score: 0})

	go func() {
		time.Sleep(time.Second * 30)
		timeUp = true
	} ()

	rand.Seed(time.Now().UTC().UnixNano())

	for i := range players {
		go players[i].play()
	}

	for timeUp == false {
		// Game in progress...
	}

	fmt.Println("Time up!")
}