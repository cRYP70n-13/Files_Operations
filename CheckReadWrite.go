package main

import (
	"log"
	"os"
)

func main() {
	// Test write permissions. It is possible the file
	// Does not exist and that will return a different
	// error that can be checked with os.IsNotExist(err)
	file, err := os.OpenFile("Test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied")
		}

		// Test the read permission
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied")
		}
	}
	defer file.Close()
}
