package main

import (
	"fmt"
	"time"
)

func main() {
	numCh := make(chan bool)
	engCh := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		numCh <- true

	}()
	time.Sleep(time.Second)

	go func() {
		i := 1
		for {
			switch {
			case <-numCh:
				fmt.Print(i, i+1)
				i += 2
				engCh <- true
			}
		}
	}()

	go func() {
		engStr := "ABCDEFGHIJK"
		i := 0
		for {

			switch {
			case <-engCh:
				fmt.Print(engStr[i : i+2])
				i += 2
				numCh <- true
			}
		}
	}()
	for {
	}
}
