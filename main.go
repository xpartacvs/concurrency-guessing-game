package main

import (
	"concurrency-guessing-game/player"
	"context"
	"log"
	"math/rand"
	"time"
)

func main() {
	min, max := 1, 1000
	logger := log.Default()
	chanGuess := make(chan int)
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := randomizer.Intn((max - min) + min)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 20; i++ {
		p := player.NewPlayer(i, min, max, number, logger, randomizer, ctx, cancel)
		go p.Start(chanGuess)
	}

	logger.Printf("THE WINNER IS PLAYER-%d", <-chanGuess)
}
