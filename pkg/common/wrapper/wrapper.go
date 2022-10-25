package wrapper

import (
	"fmt"
	"runtime/debug"
	"sync"
)

type Wrapper struct {
	sync.WaitGroup
}

func (w *Wrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Wrapper: %v", err)
				debug.PrintStack()
			}
		}()
		cb()
		w.Done()
	}()
}
