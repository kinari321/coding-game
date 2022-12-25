package main

import (
	"fmt"
	"strconv"
)

func main() {
	if "3/56" == PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz") {
		fmt.Println("正解やで")
	} else {
		fmt.Println("まちごうとるやで")
	}
	if "6/60" == PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz") {
		fmt.Println("正解やで")
	} else {
		fmt.Println("まちごうとるやで")
	}
}

func PrinterError(s string) string {
	count := 0
	for _, c := range s {
		switch c {
		case 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			count++
		}
	}
	return strconv.Itoa(count) + "/" + strconv.Itoa(len(s))
}

// Other Answer 1
func PrinterError1(s string) string {
	e := 0
	for _, c := range s {
		if c >= 'a' && c <= 'm' {
			continue
		} else {
			e++
		}
	}
	return fmt.Sprintf("%d/%d", e, len(s))
}

// Other Answer 2
func PrinterError2(s string) string {
	errors := 0
	for _, c := range s {
		if string(c) > "m" {
			errors++
		}
	}
	return fmt.Sprintf("%d/%d", errors, len(s))
}
