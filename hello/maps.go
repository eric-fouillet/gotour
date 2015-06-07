package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

type School struct {
	City, Name string
}

func _2_main() {
	//m := make(map[string]School)
	var m = map[string]School{"john": {"Boston", "MIT"}}
	//m["john"] = School{"Toyota", "Prius"}
	fmt.Println(m)
	findInMap("john", m)
	delete(m, "john")
	findInMap("john", m)
	wc.Test(WordCount)
}

func findInMap(key string, m map[string]School) {
	v, ok := m[key]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Printf("key %v not found\n", key)
	}
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for i := 0; i < len(words); i++ {
		result[words[i]] += 1
	}
	return result
}
