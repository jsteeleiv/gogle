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

func (tree *Btree) Insert(key, value string){

}
