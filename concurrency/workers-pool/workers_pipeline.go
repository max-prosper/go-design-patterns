package main

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: s,
		Handler: func(i interface{}) {
			defer wg.Done()
			str, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(str)
		},
	}
	return myRequest
}

func main() {
	bufferSize := 100
	dispatcher := NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{
			id:      i,
			prefixS: fmt.Sprintf("WorkerID: %d ->", i),
			sufixS:  " World",
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()

	wg.Wait()
}
