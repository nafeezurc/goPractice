// Goroutines

package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex futher helps with concurrency by helping avoid threads trying to write to the same memory location
// Lock() and Unlock() methods which pauses a thread's operation until unlocked
// var m = sync.Mutex{}
// read and write mutex, includes RLock() and RUnlock() methods on top of Lock() and Unlock()
var m = sync.RWMutex{}

// wait group to deal with concurrency
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	// using the time library, we can see how much time executing this takes
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// running this normally (sequentially), we see it takes nearly 4s to complete
		// using the go keyword, we can run this concurrently, meaning it'll keep moving on in the next step without waiting for one call to finish
		// without a wait group though, the tasks may not be completed at all and the program will just move on
		// add 1 to wait group whenever running concurrent operations
		wg.Add(1)
		go dbCall(i)
	}
	// add a wg.Wait() to make sure the program waits for execution completion
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	// setting a lock will make the program wait until a lock is released
	/*m.Lock()
	results = append(results, dbData[i])
	m.Unlock()*/

	// using a RWMutex
	save(dbData[i])
	// when a Goroutine reaches log(), it checks if there is a full lock
	// meaning it checks if theres a Lock(), meaning we can't read while results is being written to
	// if there is none, it'll acquire a RLock() before proceeding with the rest of the code
	log()
	// let program know execution complete, this decrements the wait group counter
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}
