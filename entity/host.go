package entity

func NewHost(permissions int) Host {
	return Host{Eating: make(chan int, permissions), DoneEating: make(chan int, permissions)}
}

type Host struct {
	Eating     chan int // buffered channels with buffer size 2 (for two philosophers to ea)
	DoneEating chan int
}

func (h *Host) AssignPermitions() {
	for {
		select {
		// wait until there will be free permission
		case <-h.DoneEating:
			// give permission
		
		<-h.Eating
		
		}

}}