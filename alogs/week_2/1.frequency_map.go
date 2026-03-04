package main

import "fmt"

func CharFrequency(s string) map[rune]int {
	m := map[rune]int{}

	for _, ch := range s {
		m[ch]++
	}

	return m
}

func FirstUniqueChar(s string) rune {
	m := make(map[rune]int, len(s))

	for _, ch := range s {
		m[ch]++
	}

	for _, letter := range s {
		if m[letter] == 1 {
			return letter
		}
	}

	return 0
}

func StreamFirstUniqueChar(ch rune, nPtr *map[rune]*Node, mPtr *map[rune]int) rune {

	m := *mPtr
	n := *nPtr

	m[ch]++

	if m[ch] == 1 {
		node := n[ch]
		node.value = ch

	} else if m[ch] > 1 {
		node := n[ch]
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	return 0
}

type Node struct {
	prev  *Node
	next  *Node
	value rune
}

type StreamUnique struct {
	freq  map[rune]int
	nodes map[rune]*Node
	head  *Node
	tail  *Node
}

type add() {

}

func main() {
	m := map[rune]int{}
	n := map[rune]*Node{}

	str := "CCASPomas🙁p5ocm12#MA😁P12SO(😒F%#39Qj😒😒vapsMVAC"

	for _, s := range str {
		result := StreamFirstUniqueChar(s, &n, &m)
		fmt.Println(result)
	}

	//result := FirstUniqueChar(str)
	//fmt.Println(result)
}
