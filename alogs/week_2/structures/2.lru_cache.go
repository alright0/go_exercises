package main

import "fmt"

type NodeLru struct {
	prev  *NodeLru
	next  *NodeLru
	key   int
	value int
}

type LruCache struct {
	capacity int
	cache    map[int]*NodeLru
	head     *NodeLru
	tail     *NodeLru
}

func (c *LruCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.moveToTail(node)

	return node.value
}

func (c *LruCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}

	node, ok := c.cache[key]
	if ok {
		node.value = value
		c.moveToTail(node)
		return
	}

	node = &NodeLru{
		key:   key,
		value: value,
	}

	if len(c.cache) == c.capacity {
		c.remove(c.head)
	}

	c.addToTail(node)
	c.cache[key] = node
}

func (c *LruCache) addToTail(node *NodeLru) {
	node.next = nil
	node.prev = c.tail

	if c.head == nil {
		c.head = node
		c.tail = node
	} else {
		c.tail.next = node
		c.tail = node
	}
}
func (c *LruCache) moveToTail(node *NodeLru) {
	if c.tail == node {
		return
	}

	if c.head == node {
		c.head = node.next
		node.next.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	c.tail.next = node
	node.prev = c.tail
	c.tail = node
	node.next = nil

}

func (c *LruCache) remove(node *NodeLru) {
	if c.head == node {
		c.head = node.next
	}

	if c.tail == node {
		c.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if c.head != nil {
		c.head.prev = nil
	}

	if c.tail != nil {
		c.tail.next = nil
	}

	node.prev = nil
	node.next = nil

	delete(c.cache, node.key)

}

func main() {
	n := LruCache{
		capacity: 3,
		cache:    make(map[int]*NodeLru),
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
