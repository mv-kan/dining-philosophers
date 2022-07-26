package main

import (
	"sync"

	"github.com/mv-kan/dining-philosophers/entity"
)

func main() {
	// init all vars
	wg := sync.WaitGroup{}
	chops := [5]entity.Chopstick{}
	host := entity.NewHost(2)
	philos := [5]entity.Philosopher{
		{Number: 1, RightChop: &chops[3], LeftChop: &chops[4], Host: host},
		{Number: 2, RightChop: &chops[0], LeftChop: &chops[1], Host: host},
		{Number: 3, RightChop: &chops[1], LeftChop: &chops[2], Host: host},
		{Number: 4, RightChop: &chops[2], LeftChop: &chops[3], Host: host},
		{Number: 5, RightChop: &chops[4], LeftChop: &chops[3], Host: host},
	}
	// run go routines
	wg.Add(len(philos))
	for _, philosopher := range philos {
		go philosopher.Eat(&wg)
	}
	go host.AssignPermissions()

	// wait and close
	wg.Wait()
	// close is for closing host go routine otherwise it will run infinitly or at least until main go routine stops
	close(host.DoneEating)
}
