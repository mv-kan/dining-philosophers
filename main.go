package main

import (
	"fmt"
	"sync"

	"github.com/mv-kan/dining-philosophers/entity"
)

func main() {
	wg := sync.WaitGroup{}
	chops := [5]entity.Chopstick{}
	host := entity.NewHost(2)
	philos := [5]entity.Philosopher{
		{Number: 1, RightChop: &chops[0], LeftChop: &chops[1], Host: host},
		{Number: 2, RightChop: &chops[1], LeftChop: &chops[2], Host: host},
		{Number: 3, RightChop: &chops[2], LeftChop: &chops[3], Host: host},
		{Number: 4, RightChop: &chops[3], LeftChop: &chops[4], Host: host},
		{Number: 5, RightChop: &chops[0], LeftChop: &chops[4], Host: host},
	}
	wg.Add(len(philos))
	for _, philosopher := range philos {
		go philosopher.Eat(&wg)
	}
	go host.AssignPermissions()
	wg.Wait()
	close(host.DoneEating)
	fmt.Println("Hello it is working")
}
