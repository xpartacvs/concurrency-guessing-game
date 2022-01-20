package player

import (
	"context"
	"log"
	"math/rand"
	"time"
)

type Player struct {
	logger     *log.Logger
	min        int
	max        int
	num        int
	index      int
	randomizer *rand.Rand
	ctx        context.Context
	ctxCancel  context.CancelFunc
}

func NewPlayer(index, min, max, num int, log *log.Logger, rnd *rand.Rand, ctx context.Context, ctxCancel context.CancelFunc) *Player {
	return &Player{
		index:      index,
		min:        min,
		max:        max,
		num:        num,
		logger:     log,
		randomizer: rnd,
		ctx:        ctx,
		ctxCancel:  ctxCancel,
	}
}

func (w *Player) Start(guessNum chan int) {
	n := 0
	for {
		n++
		select {
		case <-w.ctx.Done():
			w.logger.Printf("Another Player has found the number. Player-%d has been failed.\n", w.index)
			return
		default:
			guessNumber := w.randomizer.Intn((w.max - w.min) + w.min)
			if w.num == guessNumber {
				w.logger.Printf("Finally on the %dx trials, Player-%d has guess a correct number : %d.\n", n, w.index, guessNumber)
				w.ctxCancel()
				guessNum <- w.index
				close(guessNum)
				return
			}
			w.logger.Printf("Player-%d has %dx failed on guessing number: %d.\n", w.index, n, guessNumber)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
