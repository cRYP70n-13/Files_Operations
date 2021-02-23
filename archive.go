package main

import (
	"archive/zip"
	"log"
	"os"
)

func errorHandler(err error) {
	log.Fatal(err)
}

func main() {
	// Create a file to write the archive buffer to
	// Could also use an in memory buffer.
	outFile, err := os.Create("test.zip")
	errorHandler(err)
	defer outFile.Close()

	// Create a zip writer on top of the file writer
	zipWriter := zip.NewWriter(outFile)

	// Add files to archive
	// We use some hard coded data to the purpose of demonstration
	// but we can also iterate through all the files
	// in a specific directory and pass the name and contents
	// of each file, or we can take data from our program
	// And write it in to the archive
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test1.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63"},
	}

	// Create and write files to the archive, which in turn
	// are getting written to the underlying writer to the
	// .zip file we created at the beginning
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		errorHandler(err)

		_, err = fileWriter.Write([]byte(file.Body))
		errorHandler(err)
	}

	// Clean our shit up
	err = zipWriter.Close()
	errorHandler(err)
}
