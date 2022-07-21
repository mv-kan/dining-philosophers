package entity

import "sync"

type Chopstick struct {
	Mu sync.Mutex
}
