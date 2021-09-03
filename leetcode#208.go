package main

type Node struct {
	Char byte
	Nodes map[byte]*Node
}

func InitNode(char byte) *Node {
	return &Node{
		Char: char,
		Nodes: make(map[byte]*Node),
	}
}

type Trie struct {
	nodes map[byte]*Node
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		nodes: make(map[byte]*Node),
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	chars := []byte(word)
	chars = append(chars, '0')
	nodes := this.nodes
	for _, c := range chars {
		node, ok := nodes[c]
		if !ok {
			node = InitNode(c)
			nodes[c] = node
		}
		nodes = node.Nodes
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	chars := []byte(word)
	chars = append(chars, '0')
	nodes := this.nodes
	for _, c := range chars {
		node, ok := nodes[c]
		if !ok {
			return false
		}
		nodes = node.Nodes
	}
	return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	chars := []byte(prefix)
	nodes := this.nodes
	for _, c := range chars {
		node, ok := nodes[c]
		if !ok {
			return false
		}
		nodes = node.Nodes
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */