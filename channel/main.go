package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			messages <- "hello"
		}()
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-messages)
	}

}
