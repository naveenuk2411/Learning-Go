package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Learning basic concurrency patterns")

	// Closures in goroutines -> Application to any functions in go, not just go routines.
	var str string
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		str = "Hello World"
	}()
	wg1.Wait()
	fmt.Println(str)

	// Passing a value to a goroutine
	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		// go func() {sendRPC(i)} will not work, because of how mutations work in go. The inner functions will refer the value of i
		// by reference due to closure, and by the time the go routines process the value of i they refer from the outer function, the
		// value of i would have been updated to some other value. (Maybe 4 or 5)
		// Output will be: Sending RPC call no 4, Sending RPC call no 5. Sending RPC call no 5, Sending RPC call no 5, Sending RPC call no 5
		go func(routineNo int) {
			defer wg2.Done()
			sendRPC(routineNo)
		}(i)
	}
	wg2.Wait()

	// Firing a periodic function / Periodic pings
	// time.Sleep(1 * time.Second)
	// fmt.Println("Started")
	// go periodic()
	// time.Sleep(5 * time.Second) // Waiting for a while just to see what periodic function does.

	// Firing a periodic function with kill signal
	var mu1 sync.Mutex
	time.Sleep(1 * time.Second)
	done := false
	fmt.Println("Started")
	go periodicWithKill(&done, &mu1)
	time.Sleep(5 * time.Second)
	// We can use channels as well to signal the done or kill signal
	mu1.Lock()
	done = true
	mu1.Unlock()
	fmt.Println("Cancelled")
	time.Sleep(3 * time.Second)

	// Mutexes
	counter := 0
	var mu2 sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			defer mu2.Unlock()
			mu2.Lock()
			counter++
		}()
	}
	time.Sleep(100 * time.Millisecond) // Waiting for some time to actually process counter variable 1000 times, before the main routine
	// happens to acquire a lock.
	// Obviosuly we can use wait groups for waiting on routines to exit.
	mu2.Lock()
	fmt.Println(counter)
	mu2.Unlock()

	// Use locks when accessing a shared variable across threads.
	// Another use case of locks is to ensure invariance, where we ensure the even though multiple threads manipulate the shared data
	// the invariant on that data must hold at any point in execution.
	var alice int = 10000
	var bob int = 10000
	var total int = alice + bob

	// Invariant: The total sum of alice and bob should always remain the same during the course of computation
	var mu3 sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			defer mu3.Unlock()
			mu3.Lock()
			alice += 1
			bob -= 1
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			defer mu3.Unlock()
			mu3.Lock()
			alice -= 1
			bob += 1
		}()
	}
	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu3.Lock()
		if alice+bob != total {
			fmt.Printf("Found deviation in our invariance execpted %d, found %d\n", total, alice+bob)
		}
		mu3.Unlock()
	}

	// Condition variables
	// Counting vote V1.0
	votes := 0
	finished := 0
	var mu4 sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			defer mu4.Unlock()
			vote := getVote()
			mu4.Lock()
			if vote {
				votes++
			}
			finished++
		}()
	}

	// Waiting for either the total attempts to finish or received sufficient number of votes
	mu4.Lock()
	for votes < 5 && finished != 10 {
		mu4.Unlock()
		mu4.Lock()
	}
	if votes >= 5 {
		fmt.Println("Received 5+ votes, you won!!")
	} else {
		fmt.Println("Unfortunately, you lost!")
	}
	mu4.Unlock()
}

func sendRPC(rpcNo int) {
	fmt.Println("Sending RPC call no", rpcNo)
}

func periodic() {
	for {
		fmt.Println("Ping")
		time.Sleep(1 * time.Second)
	}
}

func periodicWithKill(done *bool, mu *sync.Mutex) {
	for {
		mu.Lock()
		if *done {
			defer mu.Unlock()
			return
		}
		mu.Unlock()
		fmt.Println("Ping")
		time.Sleep(1 * time.Second)
	}
}

func getVote() bool {
	return rand.Intn(2) == 1
}
