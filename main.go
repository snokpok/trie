package main

type TrieNode struct {
	Children [26]*TrieNode
	EOW bool
}

func (n *TrieNode) CountChildren() int {
	cnt := 0
	for i := 0; i < 26; i++ {
		if n.Children[i] != nil {
			cnt++
		}
	}
	return cnt
}

type Trie struct {
	Root_ *TrieNode
}

func NewTrieNode() *TrieNode {
	newNode := TrieNode{}
	for i := 0; i < 26; i++ {
		newNode.Children[i] = nil
	}
	newNode.EOW = false
	return &newNode
}

func NewTrie() *Trie {
	return &Trie{
		Root_: NewTrieNode(),
	}
}

// find the word in this trie; if k represents a prefix then return nil since not in dictionary
func (t *Trie) Find(k string) *TrieNode {
	if t.Root_ == nil {
		t.Root_ = NewTrieNode()
		t.Root_.EOW = true
	}
	curr := t.Root_
	for i := 0; i < len(k); i++ {
		index := int(k[i]) - int('a')
		if curr.Children[index] == nil {
			return nil
		}
		curr = curr.Children[index]
	}
	if curr.EOW {
		return curr
	}
	return nil
}

func (t *Trie) Insert(k string) *TrieNode {
	if t.Root_ == nil {
		t.Root_ = NewTrieNode()
		t.Root_.EOW = true
	}
	curr := t.Root_
	for i := 0; i < len(k); i++ {
		index := int(k[i]) - int('a')
		if curr.Children[index] == nil {
			curr.Children[index] = NewTrieNode()
		}
		curr = curr.Children[index]
	}
	curr.EOW = true
	return curr
}

// remove the node with key k in the trie; if there's no such node with k then don't do anything
func (t *Trie) Remove(k string) {
	if t.Root_ == nil {
		return
	}
	curr := t.Root_
	t.helperRemoveBubbleUp(curr, k, 0)
}

// from curr, going upward removing each node until node has >1 children or is a word
func (t *Trie) helperRemoveBubbleUp(curr *TrieNode, k string, lvl int) *TrieNode {
	if curr == nil {
		return nil
	}
	if lvl == len(k) {
		if curr.EOW {
			curr.EOW = false
		}
		if curr.CountChildren() == 0 {
			curr = nil
		}
		return curr
	}
	index := int(k[lvl]) - int('a')
	curr.Children[index] = t.helperRemoveBubbleUp(curr.Children[index], k, lvl+1)
	if curr.CountChildren() == 0 && !curr.EOW {
		curr = nil
	}
	return curr
}

func main() {
	t := NewTrie()
	t.Insert("search")
}
