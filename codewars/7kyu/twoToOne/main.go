package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "aretheyhere"
	s2 := "yestheyarehere"
	fmt.Println(TwoToOne(s1, s2))
}

// My Answer
func TwoToOne(s1 string, s2 string) string {
	concatArray := strings.Split(s1+s2, "")
	sort.Strings(concatArray)
	result := []string{}
	m := make(map[string]bool)
	for _, v := range concatArray {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return strings.Join(result, "")
}
