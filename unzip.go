package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create a reader out of the zip archive
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// Iterate over each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFiles, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFiles.Close()

		// Specifies what the extracted file name should be.
		// you can specify a full path or a prefix
		// To move it to a different directory.
		// In this case, we will extract the file from
		// The zip to a file of the same name
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating Directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFiles)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
