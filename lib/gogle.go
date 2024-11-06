package lib

import(
	"flag"
	"fmt"
	"os"
	//"strings"
)

func ParseArgs() (string, string, string){
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

func BuildJdb() (db){
	return db
}
