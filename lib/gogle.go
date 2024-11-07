package lib

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ParseArgs() (string, string, string) {
	//define command-line args
	search_term := flag.String("s", "", "Search Term")
	search_path := flag.String("p", ".", "Directory path to search")
	search_thrd := flag.String("t", "1", "Threads to utilize in search")
	flag.Parse()

	// check if search term provided
	if *search_term == "" {
		fmt.Println("Please provide a search term with the `-s` flag.")
		os.Exit(1)
	}

	return *search_term, *search_path, *search_thrd
}

func PopulateBtree(root string, btree *Btree) {
	//initialize map
	dirs := make(map[string]*Directory)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error: accessing path %s: %v", path, err)
			return err
		}
		if info.IsDir() {
			// directory creation
			dir := &Directory{
				Path:    path,
				Files:   []*FileData{},
				Subdirs: []*Directory{},
			}
			dirs[path] = dir
			// add to Btree
			btree.Insert(path)
		} else {
			// file creation
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				log.Printf("Error: reading file %s: %v", path, err)
				return readErr
			}
			file := &FileData{
				Path:     path,
				Data:     data,
				Size:     info.Size(),
				Modified: true,
				Mod_time: info.ModTime()}
			parentDir := filepath.Dir(path)
			if dir, exists := dirs[parentDir]; exists {
				dir.Files = append(dir.Files, file)
			}
			// add to btree
			btree.Insert(path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk Directory %s: %v", root, err)
	}
}

func PrintBtree(btree *Btree) {
	// define recursive function
	var printNode func(node *BtreeNode, level int)
	printNode = func(node *BtreeNode, level int) {
		if node == nil {
			return
		}

		// variable to easily adjust indentation for each level
		indent := ""
		for i := 0; i < level; i++ {
			indent += "  "
		}
		// print each node entry
		for _, key := range node.keys{
			fmt.Printf("%sKey: %s\n", indent, key)

			// switch data := key.value.(type) {
			// case *FileData:
			// 	fmt.Printf("%sFile: %s, Size: %d bytes, Modified: %d\n",
			// 		indent, data.Path, data.Size, data.Mod_time)
			// case *Directory:
			// 	fmt.Printf("%sDirectory: %s\n", indent, data.Path)
			// default:
			// 	fmt.Printf("%sError: Unknown Type ...", data.Path)
			// }
		}

		//recursively print child nodes
		for _, child := range node.children {
			printNode(child, level+1)
		}
	}
	// start printing from root
	printNode(btree.root, 0)
}
