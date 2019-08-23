package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const sentence = "The cat sat on the mat. The mat was blue."

func main() {
	reg := regexp.MustCompile("\\W+")
	words := reg.Split(sentence, -1)
	counts := make(map[string]int)

	for _, word := range words {
		if len(word) > 0 {
			counts[strings.ToUpper(word)]++
		}
	}

	sortedTwins := sortTwinList(counts)

	for _, twin := range sortedTwins {
		fmt.Printf("%v : %d\n", twin.Key, twin.Value)
	}
}

type twin struct {
	Key   string
	Value int
}

type twinList []twin

func (t twinList) Len() int {
	return len(t)
}

func (t twinList) Less(i, j int) bool {
	return t[i].Key < t[j].Key
}

func (t twinList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func sortTwinList(m map[string]int) twinList {
	tl := make(twinList, len(m))
	i := 0
	for k, v := range m {
		tl[i] = twin{k, v}
		i++
	}
	sort.Sort(tl)
	return tl
}
