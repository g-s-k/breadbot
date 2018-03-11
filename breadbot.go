package main

import "fmt"
import "time"

//import "image"

const searchInterval = 5

func main() {
	// pipelines
	c1 := make(chan string)
	go initiate(c1)
	go imgGet(c1)
	// separate goroutines for different functions
	// close out program
	var input string
	fmt.Scanln(&input)
}

func initiate(ch chan<- string) {
	for {
		ch <- "foo"
		time.Sleep(time.Second * searchInterval)
	}
}

func imgGet(cin chan string) {
	for {
		msg := <-cin
		fmt.Println(msg)
		time.Sleep(time.Second * searchInterval)
	}
}
