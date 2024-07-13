package waitgroup_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/goaux/waitgroup"
)

func Example() {
	var sy waitgroup.Sync
	results := make(chan string, 3)

	sy.Go(func() {
		time.Sleep(100 * time.Millisecond)
		results <- "First task done"
	})

	sy.Go(func() {
		time.Sleep(200 * time.Millisecond)
		results <- "Second task done"
	})

	sy.Go(func() {
		time.Sleep(50 * time.Millisecond)
		results <- "Third task done"
	})

	// Wait for all goroutines to complete
	sy.Wait()
	close(results)

	// Print results in the order they were added to the channel
	for result := range results {
		fmt.Println(result)
	}
	// Output:
	// Third task done
	// First task done
	// Second task done
}

func TestSync_Go(t *testing.T) {
	t.Run("Single goroutine", func(t *testing.T) {
		var sy waitgroup.Sync
		var counter int32

		sy.Go(func() {
			atomic.AddInt32(&counter, 1)
		})

		sy.Wait()

		if atomic.LoadInt32(&counter) != 1 {
			t.Errorf("Expected counter to be 1, got %d", counter)
		}
	})

	t.Run("Multiple goroutines", func(t *testing.T) {
		var sy waitgroup.Sync
		var counter int32
		numGoroutines := 10

		for i := 0; i < numGoroutines; i++ {
			sy.Go(func() {
				time.Sleep(10 * time.Millisecond) // Simulate some work
				atomic.AddInt32(&counter, 1)
			})
		}

		sy.Wait()

		if atomic.LoadInt32(&counter) != int32(numGoroutines) {
			t.Errorf("Expected counter to be %d, got %d", numGoroutines, counter)
		}
	})

	t.Run("Nested goroutines", func(t *testing.T) {
		var sy waitgroup.Sync
		var counter int32

		sy.Go(func() {
			atomic.AddInt32(&counter, 1)
			sy.Go(func() {
				atomic.AddInt32(&counter, 1)
			})
		})

		sy.Wait()

		if atomic.LoadInt32(&counter) != 2 {
			t.Errorf("Expected counter to be 2, got %d", counter)
		}
	})
}
