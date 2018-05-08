package main

import "fmt"

var ch = make(chan bool)

func main() {
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- true
}
