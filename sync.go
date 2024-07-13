// Package waitgroup provides a wrapper around sync.WaitGroup
// to simplify launching goroutines and waiting for their completion.
package waitgroup

import "sync"

// Sync is a wrapper struct around sync.WaitGroup that provides
// a convenient method for launching goroutines.
type Sync struct {
	sync.WaitGroup
}

// Go launches the given function in a new goroutine and automatically
// handles the addition and completion signaling for the wait group.
// This method simplifies the process of using WaitGroup with goroutines.
func (sy *Sync) Go(fn func()) {
	sy.Add(1)
	go func() {
		defer sy.Done()
		fn()
	}()
}
