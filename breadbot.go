package main

import "fmt"

//import "image"

func main() {
	// pipelines
	c1 := make(chan string)
	go imgGet(c1)
	c1.send("blah")
	// separate goroutines for different functions
	// close all the channels
	close(c1)
}

func imgGet(ch chan string) {
	for ch.recv() {
		fmt.Println(ch)
	}
}
