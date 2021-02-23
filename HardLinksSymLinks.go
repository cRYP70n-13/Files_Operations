package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a hard link
	// we will have two file names that point to the same contents
	// Changing the content of one will change to the other
	// Deleting/renaming one will not affect the other
	err := os.Link("Original.txt", "Original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lstat will return file info, but if it is actually
	// a symlink, it will return info about the symlink.
	// It will not follow the link and give information
	// About the real life
	// NOTE: Symlink do not work in Windows
	fileInfo, err := os.Lstat("Original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Link info: %+v", fileInfo)

	// Change ownership of a symlink only and not the file it points to
	err = os.Lchown("Original_sym.txt", os.Getuid(), os.Getegid())
	if err != nil {
		log.Fatal(err)
	}
}
