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

// 1. Other Answer
func TwoToOne1(s1 string, s2 string) string {
	chars := strings.Split(s1+s2, "")
	sort.Strings(chars)
	result := ""
	for _, s := range chars {
		chr := string(s)
		if !strings.Contains(result, chr) {
			result = result + chr
		}
	}
	return result
}

// 2. Other Answer
func TwoToOne2(s1 string, s2 string) string {
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s1+s2, char) {
			result += char
		}
	}
	return result
}
