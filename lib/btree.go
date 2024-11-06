package jdb

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type BtreeNode struct {
	keys []string
	children []*BtreeNode
	isLeaf bool
}

type Btree struct {
	root *BtreeNode
	branches int
}

func NewBtree(branches int) *Btree {
	return &Btree{
		root: &BtreeNode{isLeaf: true},
		branches: branches,
	}
}

func (tree *Btree) splitChild(parent *BtreeNode, index int){
	node := parent.children[index]
	mid := len(node.keys) / 2
	
	newNode := &BtreeNode{isLeaf: node.isLeaf}
	newNode.keys = append(newNode.keys, node.keys[mid+1:]...)
	parent.keys = append(
		parent.keys[:index],
		append([]string{node.keys[mid]}, parent.keys[index:]...)...
	)
	parent.children = append(
		parent.children[:index+1],
		append([]*BtreeNode{newNode}, parent.children[index+1:]...)...
	)
	node.keys = node.keys[:mid]
}

// TODO: implement insertNonFull func

// TODO: implement save to/load from disk

func (tree *Btree) Insert(key, value string){
	root := tree.root
	if len(root.keys) == tree.branches-1{
		newRoot := &BtreeNode{isLeaf: false}
		newRoot.children = append(newRoot.children, root)
		tree.root = newRoot
		tree.splitChild(newRoot, 0)
	}
	tree.insertNonFull(root, key, value)
}


