package main

import "time"

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func NewDispatcher(b int) *dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}
