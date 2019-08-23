package main

import (
	"fmt"
	"regexp"
	"strings"
)

const sentence = "The cat sat on the mat. The mat was blue."

func main() {
	reg := regexp.MustCompile("\\W+")
	words := reg.Split(sentence, -1)
	dict := make(map[string]int)

	for _, word := range words {
		if len(word) > 0 {
			dict[strings.ToUpper(word)]++
		}
	}

	for word, count := range dict {
		fmt.Printf("%v : %d\n", word, count)
	}
}
