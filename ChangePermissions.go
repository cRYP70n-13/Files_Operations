package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// Change permissions using linux style
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change the ownership
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change timestamps
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifiedTime := twoDaysFromNow

	err = os.Chtimes("test.txt", lastAccessTime, lastModifiedTime)
	if err != nil {
		log.Println(err)
	}
}
