package main

import "fmt"

func StreamFirstUniqueChar(s string, lst *StreamUnique) rune {
	for _, ch := range s {
		lst.freq[ch]++

		if lst.freq[ch] == 1 {
			lst.Add(ch)
		} else if lst.freq[ch] == 2 {
			lst.Remove(ch)
		}
		fmt.Println(lst.First())
	}
	return 0
}

type Node_ struct {
	prev  *Node_
	next  *Node_
	value rune
}

type StreamUnique struct {
	freq  map[rune]int
	nodes map[rune]*Node_
	head  *Node_
	tail  *Node_
}

func (s *StreamUnique) Add(ch rune) {
	node := &Node_{
		value: ch,
	}

	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.next = node
		node.prev = s.tail
		s.tail = node
	}

	s.nodes[ch] = node

}

func (s *StreamUnique) Remove(ch rune) {
	node, ok := s.nodes[ch]
	if !ok {
		return
	}

	if s.head == node {
		s.head = node.next
	}

	if s.tail == node {
		s.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if s.head != nil {
		s.head.prev = nil
	}

	if s.tail != nil {
		s.tail.next = nil
	}

	node.prev = nil
	node.next = nil

}

func (s *StreamUnique) First() rune {
	if s.head == nil {
		return 0
	}
	return s.head.value
}

func main() {
	n := StreamUnique{
		freq:  make(map[rune]int),
		nodes: make(map[rune]*Node_),
		head:  nil,
		tail:  nil,
	}

	str := "abcabddabce"

	StreamFirstUniqueChar(str, &n)

}
