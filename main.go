package main

import (
	"fmt"
	"gogle/lib"
)

func main(){
	fmt.Println("Welcome to Gogle!")
	// add functionality here
	// term, path, threads := lib.ParseArgs()

	//initialize btree for storing files
	db := lib.NewBtree(10)
	path := "./tests"
	lib.PopulateBtree(path, db)
        // Now `btree` is populated with all files and their contents
	// fmt.Println(db)
	lib.PrintBtree(db)
}
