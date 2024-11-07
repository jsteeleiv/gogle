package lib

type TrieNode struct{
	children map[rune]*TrieNode // https://go.dev/blog/strings
	isEndofWord bool
}

type Trie struct{
	root *TrieNode
}

func NewTrie() *Trie{
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (trie *Trie) Insert(word string){
	node := trie.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEndofWord = true
}

func (trie *Trie) Search(word string) bool {
node := trie.root
for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEndofWord
}

// check if any word in trie that starts with prefix
func (trie *Trie) StartsWith(prefix string) bool {
	node := trie.root
for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}
