package main

import (
	"fmt"
	"strings"
)

func main() {
	if 3 == FindShort("turns out random test cases are easier than writing out basic ones") {
		fmt.Println("正解やで")
	} else {
		fmt.Println("まちごうとるやで")
	}
	if 2 == FindShort("Let's travel abroad shall we") {
		fmt.Println("正解やで")
	} else {
		fmt.Println("まちごうとるやで")
	}
}

func FindShort(s string) int {
	array := strings.Split(s, " ")
	shortestWord := len(array[0])
	for _, v := range array {
		if len(v) < shortestWord {
			shortestWord = len(v)
		}
	}
	fmt.Println(shortestWord)
	return shortestWord
}

// Other Answer 1
func FindShort1(s string) int {
	shortest := len(s)
	for _, word := range strings.Split(s, " ") {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest
}

// Other Answer 2
func FindShort2(s string) int {
	var t int
	for i, s := range strings.Fields(s) {
		if len(s) < t || i == 0 {
			t = len(s)
		}
	}
	return t
}
