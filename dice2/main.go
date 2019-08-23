package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	faces    = []int{1, 2, 3, 4, 5, 6}
	onlyOnce sync.Once
)

func rollDice() int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	return faces[rand.Intn(len(faces))]
}

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", rollDice())
	}
	fmt.Println()
}
