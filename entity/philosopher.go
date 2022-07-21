package entity

import "fmt"

type Philosopher struct {
	RightChop *Chopstick
	LeftChop  *Chopstick
	Number    int
}

func (ph *Philosopher) Eat(eating, doneEating chan int) {
	// ask for permission
	eating <- ph.Number

	ph.RightChop.Mu.Lock()
	ph.LeftChop.Mu.Lock()

	fmt.Println("starting to eat ", ph.Number)
	fmt.Println("finished eating ", ph.Number)

	ph.RightChop.Mu.Unlock()
	ph.LeftChop.Mu.Unlock()

	// tell host that this philosopher done eating
	doneEating <- ph.Number
}
