package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	one = `

  #
`
	two = `
    #

#`
	three = `
    #
  #
#`
	four = `
#   #

#   #`
	five = `
#   #
  #
#   #`
	six = `
#   #
#   #
#   #`
	faces    = []string{one, two, three, four, five, six}
	onlyOnce sync.Once
)

func rollDice() int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	return rand.Intn(len(faces))
}

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("-----%s\n-----\n", faces[rollDice()])
	}
	fmt.Println()
}
