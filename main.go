package main

import (
	"fmt"
	//"os"
	//"path/filepath"
	"gogle/lib"
)

func main(){
	fmt.Println("Welcome to Gogle!")
	// add functionality here
	term, path, threads := lib.ParseArgs()
	fmt.Println(term, path, threads)

}
