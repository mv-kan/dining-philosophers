package entity

import "fmt"

func NewHost(permissions int) Host {
	return Host{Eating: make(chan int, permissions), DoneEating: make(chan int, permissions)}
}

type Host struct {
	Eating     chan int // buffered channels with buffer size 2 (for two philosophers to ea)
	DoneEating chan int
}

func (h *Host) AssignPermissions() {
	// get info that a philo done with eating and ...
	// I close DoneEating chan in main
	for number := range h.DoneEating {
		// ... empty eating chan to give permission to eat for another philo
		fmt.Printf("Host: Philosopher number %d done eating\n", number)
		<-h.Eating
	}
}
