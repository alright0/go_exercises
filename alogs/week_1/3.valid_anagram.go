package main

import "fmt"

func main() {
	str1 := "ADADALDSLV"
	str2 := "VADADALDSL"

	result := isAnagram(str1, str2)
	fmt.Println(result)
}

func isAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	aMap := make(map[rune]int)

	for _, ch := range str1 {
		if _, ok := aMap[ch]; ok {
			aMap[ch]++
		} else {
			aMap[ch] = 1
		}
	}

	for _, ch := range str2 {
		if _, ok := aMap[ch]; ok {
			aMap[ch]--
		} else {
			return false
		}
	}

	for _, v := range aMap {
		if v != 0 {
			return false
		}
	}
	return true
}
