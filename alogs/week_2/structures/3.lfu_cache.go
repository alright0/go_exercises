package main

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
	freq  int
}

type List struct {
	head *Node
	tail *Node
}

func (list *List) add(node *Node) {
	prev := list.tail.prev

	prev.next = node
	node.prev = prev

	node.next = list.tail
	list.tail.prev = node
}

func (list *List) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil
}

func (list *List) removeFirst() *Node {
	node := list.head.next

	if node == list.tail {
		return nil
	}

	list.remove(node)

	return node
}

func NewList() *List {
	headNode := &Node{}
	tailNode := &Node{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &List{
		head: headNode,
		tail: tailNode,
	}
}

type LfuCache struct {
	capacity int
	size     int
	minFreq  int

	keyMap  map[int]*Node
	freqMap map[int]*List
}

func (c *LfuCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}

	node.freq++

	c.moveToTail(node)

	return node.value
}

func (c *LfuCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}

	node, ok := c.cache[key]
	if ok {
		node.value = value
		c.moveToTail(node)
		return
	}

	node = &Node{
		key:   key,
		value: value,
	}

	if len(c.cache) == c.capacity {
		c.remove(c.head)
	}

	c.addToTail(node)
	c.cache[key] = node
}

func Constructor(capacity int) LfuCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return LfuCache{
		capacity: capacity,
		size:     0,
		minFreq:  0,
		keyMap:   make(map[int]*Node),
		freqMap:  make(map[int]*List),
	}
}

func (c *LfuCache) add(node *Node)    {}
func (c *LfuCache) remove(node *Node) {}
func (c *LfuCache) removeLRU() *Node  {}

func (c *LfuCache) increaseFreq(node *Node) {}

//func (c *LfuCache) addToTail(node *Node) {
//	node.next = nil
//	node.prev = c.tail
//
//	if c.head == nil {
//		c.head = node
//		c.tail = node
//	} else {
//		c.tail.next = node
//		c.tail = node
//	}
//}
//func (c *LfuCache) moveToTail(node *Node) {
//	if c.tail == node {
//		return
//	}
//
//	if c.head == node {
//		c.head = node.next
//		node.next.prev = nil
//	} else {
//		node.prev.next = node.next
//		node.next.prev = node.prev
//	}
//
//	c.tail.next = node
//	node.prev = c.tail
//	c.tail = node
//	node.next = nil
//
//}
//
//func (c *LfuCache) remove(node *Node) {
//	if c.head == node {
//		c.head = node.next
//	}
//
//	if c.tail == node {
//		c.tail = node.prev
//	}
//
//	if node.prev != nil {
//		node.prev.next = node.next
//	}
//
//	if node.next != nil {
//		node.next.prev = node.prev
//	}
//
//	if c.head != nil {
//		c.head.prev = nil
//	}
//
//	if c.tail != nil {
//		c.tail.next = nil
//	}
//
//	node.prev = nil
//	node.next = nil
//
//	delete(c.cache, node.key)
//
//}

func main() {
	n := LfuCache{
		capacity: 3,
		cache:    make(map[int]*Node),
		head:     nil,
		tail:     nil,
	}

	n.Put(1, 1)
	n.Put(2, 2)
	n.Put(3, 3)
	n.Put(4, 4)
	val := n.Get(1)
	fmt.Println("val:", val)

	val = n.Get(2)
	fmt.Println("val:", val)
}
