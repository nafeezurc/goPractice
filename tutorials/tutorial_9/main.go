// Channels

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		// to create a channel, use the make function
		// generally, consider a channel as containing an underlying array
		var c = make(chan int)

		// using a goroutine to add a value to the channel
		go process(c)

		// can retrieve a value from a channel using similar syntax
		// <-c
		// this pops the value from the channel, and the channel becomes empty

		// using loops with populated channel
		for val := range c {
			fmt.Println(val)
		}
		// the go routine will wait for main to pop the value from the channel before continuing with the process and adding more values
		// again due to unbuffered channels only being able to store one value
		// buffered channels on the other hand, can store multiple values
	*/ // This comment block contains unbuffered channel tutorial

	// to make a buffered channel, add a size parameter as shown
	var c = make(chan int, 5)

	// we'll do the same process again
	// note that we exit the process before main finishes executing this time
	// the opposite was the case with the unbuffered channel
	go process(c)
	for val := range c {
		fmt.Println(val)
		time.Sleep(time.Second * 1)
	}

	// running the chicken example
	CheckChicken()
}

func process(c chan int) {
	// use the defer keyword to close the channel to close it before the function returns
	defer close(c)
	// this is important to let the program know when to stop an infinite reading loop
	// otherwise you get a deadlock error
	// alternatively just add a close statement at the end without the defer

	// channels have special syntax
	// use an arrow like this <- to add to a channel
	c <- 123
	// this is an unbuffered channel, which can only have one value
	// another thing to consider is we can't directly write to an unbuffered channel without reading from it first
	// this is where goroutines come in

	// using loops to populate channels
	for i := 0; i < 5; i++ {
		c <- i
	}

	fmt.Println("Exiting process")
}
