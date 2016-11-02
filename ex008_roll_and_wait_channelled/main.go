package main

import (
	"fmt"
	"time"
	"math/rand"
)

type person struct {
	name  string
	score int
}

func (player *person) play() {
	for {
		roll := rand.Intn(6) + 1
		wait := 6 - roll

		if (wait == 0) {
			fmt.Printf("%s rolled a %d, rolling again immediately\n", player.name, roll)
		} else {
			fmt.Printf("%s rolled a %d, waiting %ds\n", player.name, roll, wait)
		}

		resultChannel <- rollResult{whoRolled: player, rolled: roll}
		time.Sleep(time.Second * time.Duration(wait))
	}
}

type rollResult struct {
	whoRolled *person
	rolled    int
}

var controlChannel = make(chan bool, 1)
var resultChannel = make(chan rollResult, 2)

func main() {
	currentWinner := &person{name: "", score: 0}

	players := []person{}
	players = append(players, person{name: "Ed", score: 0})
	players = append(players, person{name: "Natasha", score: 0})

	rand.Seed(time.Now().UTC().UnixNano())

	// Game termination timer
	go func() {
		time.Sleep(time.Second * 30)
		controlChannel <- true
	}()

	// Start the players rolling
	for i := range players {
		go players[i].play()
	}

	GameLoop:
	for {
		select {
		case newRoll := <-resultChannel:
			newRoll.whoRolled.score = newRoll.whoRolled.score + newRoll.rolled

			if (newRoll.whoRolled.score > currentWinner.score) {
				currentWinner = newRoll.whoRolled
			}
		case <-controlChannel:
			fmt.Println("\nTime up!")

			for i := range players {
				fmt.Printf("%s scored %d\n", players[i].name, players[i].score)
			}

			fmt.Printf("%s wins with a score of %d\n", currentWinner.name, currentWinner.score)
			break GameLoop
		}
	}
}