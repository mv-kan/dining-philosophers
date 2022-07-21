package entity

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	RightChop *Chopstick
	LeftChop  *Chopstick
	Number    int
	Host      Host
}

// IMPORTANT THING
// DO NOT MAKE THIS INSTANCE AS POINTER
// it can lead to unpredicted behaviour especially if you relying on
// fields in struct
func (ph Philosopher) Eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		// ask for permission
		ph.Host.Eating <- ph.Number

		ph.RightChop.Mu.Lock()
		ph.LeftChop.Mu.Lock()

		fmt.Println("starting to eat ", ph.Number)
		fmt.Println("finished eating ", ph.Number)

		ph.LeftChop.Mu.Unlock()
		ph.RightChop.Mu.Unlock()

		// tell host that this philosopher done eating
		ph.Host.DoneEating <- ph.Number
	}
	wg.Done()
}
