package main

import "fmt"
import "time"

//import "image"

const searchInterval = 5
const imgApiAddress = "https://www.googleapis.com/customsearch/v1"

func main() {
	// pipelines
	c1 := make(chan string)
	go initiate(c1)
	go imgGet(c1)
	// separate goroutines for different functions
	// close out program
	for {
	}
}

func initiate(ch chan<- string) {
	for {
		var input string
		fmt.Scanln(&input)
		ch <- input
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
