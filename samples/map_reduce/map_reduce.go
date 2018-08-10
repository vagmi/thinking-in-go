package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START_GENERATOR_OMIT
func randomDurations() <-chan time.Duration {
	outChan := make(chan time.Duration, 1)
	go func() {
		for i := 0; i < 10; i++ {
			dur := time.Second * time.Duration(rand.Intn(10)+1)
			fmt.Printf("generator: Going to sleep for %d seconds\n", int(dur.Seconds()))
			outChan <- dur
		}
		close(outChan)
	}()
	return outChan
}

//END_GENERATOR_OMIT

// START_WORKER_OMIT
func worker(id int, durations <-chan time.Duration) <-chan int {
	outChan := make(chan int)
	go func() {
		for dur := range durations {
			afterChan := time.After(dur)
			fmt.Printf("Worker %d: waiting for %d seconds\n", id, int(dur.Seconds()))
			<-afterChan
			outChan <- int(dur.Seconds())
		}
		close(outChan)
	}()
	return outChan
}

// END_WORKER_OMIT

//START_REDUCE_OMIT
func reduce(chan1, chan2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		timer := time.NewTimer(10 * time.Second)
	Outer:
		for {
			select {
			case dur1, ok := <-chan1:
				if ok {
					c <- dur1
					timer.Reset(10 * time.Second)
				}
			case dur2, ok := <-chan2:
				if ok {
					c <- dur2
					timer.Reset(10 * time.Second)
				}
			case <-timer.C:
				break Outer
			}
		}
		close(c)
	}()
	return c
}

//END_REDUCE_OMIT

// START_REDUCED_OMIT
func reduceDynamic(chans []<-chan int) <-chan int {
	outChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, c := range chans {
		go func(c <-chan int) {
			for dur := range c {
				outChan <- dur
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(outChan)
	}()
	return outChan
}

// END_REDUCED_OMIT

// START_DYNCON_OMIT
func dynamicController(wCount int) {
	data := randomDurations()
	var chans [](<-chan int)
	for i := 0; i < wCount; i++ {
		chans = append(chans, worker(i+1, data))
	}
	reduced := reduceDynamic(chans)
	for sleptFor := range reduced {
		fmt.Printf("Reduced: Slept For %d seconds\n", sleptFor)
	}
}

// END_DYNCON_OMIT

// START_CONTROLLER_OMIT
func controller() {
	data := randomDurations()
	// map
	w1 := worker(1, data)
	w2 := worker(2, data)
	// reduce
	reduced := reduce(w1, w2)
	for sleptFor := range reduced {
		fmt.Printf("Reduced: Slept For %d seconds\n", sleptFor)
	}
}

// END_CONTROLLER_OMIT
// START_MAIN_OMIT
func main() {
	start := time.Now()
	controller()
	// dynamicController(5)
	end := time.Now()
	fmt.Printf("done in %f seconds!", end.Sub(start).Seconds())
}

// END_MAIN_OMIT
