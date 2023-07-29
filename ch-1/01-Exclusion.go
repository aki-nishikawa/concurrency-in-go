package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var dataAccess sync.Mutex
	data := 0
	go func() {
		dataAccess.Lock()
		data++
		dataAccess.Unlock()
	}()

	time.Sleep(1)

	dataAccess.Lock()
	fmt.Printf("the value is %v .\n", data)
	dataAccess.Unlock()
}
