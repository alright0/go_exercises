package ds

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

func (list *List) IsEmpty() bool {
	return list.head.next == list.tail
}

type LfuCache struct {
	capacity int
	size     int
	minFreq  int

	keyMap  map[int]*Node
	freqMap map[int]*List
}

func (c *LfuCache) increaseFreq(node *Node) {
	oldFreq := node.freq

	oldList := c.freqMap[oldFreq]
	oldList.remove(node)

	if oldFreq == c.minFreq && oldList.IsEmpty() {
		c.minFreq++
	}
	node.freq++

	newList, ok := c.freqMap[node.freq]
	if !ok {
		newList = NewList()
		c.freqMap[node.freq] = newList
	}
	newList.add(node)
}

func (c *LfuCache) Get(key int) int {
	node, ok := c.keyMap[key]
	if !ok {
		return -1
	}
	c.increaseFreq(node)
	return node.value
}

func (c *LfuCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}
	node, ok := c.keyMap[key]
	if ok {
		node.value = value
		c.increaseFreq(node)
		return
	}

	if c.size == c.capacity {
		c.removeLRU()
	}

	node = &Node{
		freq:  1,
		key:   key,
		value: value,
	}

	c.add(node)
	c.keyMap[key] = node
}

func (c *LfuCache) add(node *Node) {
	node.freq = 1
	oneFreqList, ok := c.freqMap[node.freq]
	if !ok {
		c.freqMap[node.freq] = NewList()
		oneFreqList = c.freqMap[node.freq]
	}

	oneFreqList.add(node)
	c.size++
	c.minFreq = 1
}

func (c *LfuCache) removeLRU() *Node {
	list := c.freqMap[c.minFreq]
	node := list.removeFirst()
	if node == nil {
		return nil
	}
	delete(c.keyMap, node.key)
	c.size--
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

func Constructor(capacity int) LfuCache {
	return LfuCache{
		capacity: capacity,
		size:     0,
		minFreq:  0,
		keyMap:   make(map[int]*Node),
		freqMap:  make(map[int]*List),
	}
}

func main() {
	lfuCache := Constructor(5)

	node := &Node{
		key:   1,
		value: 1,
	}

	lfuCache.add(node)
}
