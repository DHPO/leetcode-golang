package main

type DataNode struct {
	Key int
	Data int
	Prev *DataNode
	Next *DataNode
}

type LRUCache struct {
	KeyMap map[int]*DataNode
	Count int
	Capacity int
	Head *DataNode
	Tail *DataNode
}


func Constructor(capacity int) LRUCache {
	head := &DataNode{}
	tail := &DataNode{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		KeyMap: make(map[int]*DataNode),
		Count: 0,
		Capacity: capacity,
		Head: head,
		Tail: tail,
	}
}


func (this *LRUCache) Get(key int) int {
	if node, ok := this.KeyMap[key]; ok {
		this.removeNode(node)
		// add
		this.addToHead(node)

		return node.Data
	}
	return -1
}

func (this *LRUCache) addToHead(node *DataNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	node.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) removeNode(node *DataNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.KeyMap[key]; ok {
		node.Data = value
		this.removeNode(node)
		this.addToHead(node)
		return
	}
	if this.Count == this.Capacity {
		node := this.Tail.Prev
		this.removeNode(node)
		delete(this.KeyMap, node.Key)
	} else {
		this.Count += 1
	}

	node := &DataNode{
		Key: key,
		Data: value, 
	}
	this.addToHead(node)
	this.KeyMap[key] = node
}
