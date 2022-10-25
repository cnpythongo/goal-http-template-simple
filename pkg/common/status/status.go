package status

import (
	"fmt"
	"sync"
)

var (
	isRunnning bool
	wg         sync.WaitGroup
)

func init() {
	isRunnning = true
}

func IsRunnning() bool {
	return isRunnning
}

func Shutdown() {
	isRunnning = false
	fmt.Println("isRunnning false")
}

func AddWaitGroup() {
	wg.Add(1)
}

func DoneWaitGroup() {
	wg.Done()
}

func WaitGroup() {
	wg.Wait()
}
