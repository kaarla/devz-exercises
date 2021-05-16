package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	newTrie := NewTrie()
	words := []string{"car", "cat", "call", "can", "dog", "door", "domino"}
	newTrie.Build(words)
	next := newTrie.GetNext("do")
	for _, r := range next {
		fmt.Println("next: ", string(r))
	}

	fmt.Println("autocomplete", newTrie.autocomplete("do"))
}

type Node struct {
	Item string
	IsWord bool
	Children map[rune]*Node
}

func NewNode() *Node{
	newNode := Node{}
	newNode.IsWord = false
	newNode.Children = map[rune]*Node{}
	return &newNode
}

type Trie struct {
	Root *Node
	Size int
}

func NewTrie() *Trie{
	trie := Trie{}
	trie.Root = NewNode()
	trie.Size = 0
	return &trie
}

func (trie *Trie) Build(words []string) {
	for _, w := range words {
		trie.insert(w)
	}
}

func (trie *Trie) insert(word string) {
	current := trie.Root
	for _, r := range word {
		_, ok := current.Children[r]
		if ok {
			current = current.Children[r]
		} else {
			current.Children[r] = NewNode()
			trie.Size++
		}
	}
	current.IsWord = true
	current.Item = word
}

func (trie *Trie) GetNext(prefix string) []rune{
	prefixes := trie.getPrefixes(prefix)
	if prefixes == nil {
		return []rune{}
	}

	keys := []rune{}
	for r, _ := range prefixes.Children {
		keys = append(keys, r)
	}
	
	return keys
}

func (trie *Trie) getPrefixes(prefix string) *Node{
	current := trie.Root
	ok := false

	for _, r := range prefix {
		current, ok = current.Children[r]
		if !ok {
			break
		}
	}
	return current
}

func (trie *Trie) autocomplete(prefix string) []string{
	words := []string{}
	prefixes := trie.getPrefixes(prefix)
	var stack stack.Stack

	i := 0
	for _, n := range prefixes.Children {
		stack.Push(n)
		i++
	}
	
	for stack.Len() > 0 {
		node, _ := stack.Pop().(*Node)
		if node.IsWord {
			words = append(words, node.Item)
		}
		for _, n := range node.Children {
			stack.Push(n)
		}
	}
	return words
}



