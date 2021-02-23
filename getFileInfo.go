package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      err
)

func main() {
	// Stat returns file info, It will be return an error
	// If there is not file.
	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("System Interface type: %T\n:", fileInfo.Sys())
}
