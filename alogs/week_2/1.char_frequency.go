package main

import "fmt"

func CharFrequency(s string) map[rune]int {
	m := map[rune]int{}

	for _, ch := range s {
		m[ch]++
	}

	return m
}

func main() {
	str := "abcabddabce"

	result := CharFrequency(str)
	fmt.Println(result)
}
