package lib

type NodeType struct {
	Directory *Directory
	FileData  *FileData
}

type BtreeNode struct {
	// type_    NodeType
	keys     []string
	children []*BtreeNode
	// value    interface{}
	isLeaf   bool
}

type Btree struct {
	root     *BtreeNode
	order int
}

func NewBtree(order int) *Btree {
	return &Btree{
		root:     &BtreeNode{isLeaf: true},
		order: order,
	}
}

func (tree *Btree) splitChild(parent *BtreeNode, index int) {
	node := parent.children[index]
	mid := len(node.keys) / 2

	newNode := &BtreeNode{isLeaf: node.isLeaf}
	newNode.keys = append(newNode.keys, node.keys[mid+1:]...)
	parent.keys = append(
		parent.keys[:index],
		append([]string{node.keys[mid]}, parent.keys[index:]...)...,
	)
	parent.children = append(
		parent.children[:index+1],
		append([]*BtreeNode{newNode}, parent.children[index+1:]...)...,
	)
	node.keys = node.keys[:mid]
}

// insert key into node that is not full
func (bt *Btree) insertNonFull(node *BtreeNode, key string) {
	if node.isLeaf {
		node.keys = append(node.keys, key)
		// sort keys, ascending
		// sort.Strings(node.keys)
	} else {
		// find child to insert into
		for i := len(node.keys) - 1; i >= 0; i-- {
			if key < node.keys[i] {
				bt.insertNonFull(node.children[i], key)
				return
			}
		}
		bt.insertNonFull(node.children[len(node.keys)], key)
	}
}

// search node recursively for a key
func (bt *Btree) searchNode(node *BtreeNode, key string) *BtreeNode {
	for i := 0; i < len(node.keys); i++ {
		if key == node.keys[i] {
			return node
		} else if key < node.keys[i] && !node.isLeaf {
			return bt.searchNode(node.children[i], key)
		}
	}
	if !node.isLeaf {
		return bt.searchNode(node.children[len(node.keys)], key)

	}
	return nil
}

// search for a key in the Btree
func (bt *Btree) Search(key string) *BtreeNode {
	return bt.searchNode(bt.root, key)
}

// TODO: implement save to/load from disk

func (tree *Btree) Insert(key string) {
	root := tree.root
	if len(root.keys) == tree.order-1 {
		newroot := &BtreeNode{isLeaf: false}
		newroot.children = append(newroot.children, root)
		tree.root = newroot
		tree.splitChild(newroot, 0)
	}
	tree.insertNonFull(root, key)
}
