package myfonctions

import (
	"log"
	"os"
)

func Read(fileName string) string {
	// Open File.
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("error opening file :", fileName)
		os.Exit(0)
	}

	defer file.Close()

	// Get file info.
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Error getting file stats:", err)
		os.Exit(0)
	}

	// Get file size.
	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)

	// Reading data.
	_, err = file.Read(data)
	if err != nil {
		log.Println("Error reading the file:", err)
		os.Exit(0)
	}
	return string(data)
}
