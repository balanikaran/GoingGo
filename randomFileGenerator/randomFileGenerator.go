package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func main() {

	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)

	// doing this so that from where ever the user runs this file, 
	// the temp file will be stored in this directory only 
	currentPath := filepath.Dir(filename)
	tempFile, err := ioutil.TempFile(currentPath, "randomGen-")
	if err != nil {
		fmt.Println("Unable to create tempfile - ", err)
		return
	}
	
	time := time.Now().Second()
	text := []byte("This is the text to be added again and again in the file.\n")

	fileSize := getFileSizeInBytes(tempFile)
	// 10000000 == 10 MB
	for fileSize < 10000000 {
		// fmt.Println(fileSize)
		if _, err := tempFile.Write(text); err != nil {
			fmt.Println("Unable to write in file...", err)
			return
		}
		fileSize = getFileSizeInBytes(tempFile)
	}

	// close the file
	if err := tempFile.Close(); err != nil {
		fmt.Println("Error closing file...", err)
		return
	}

	os.Rename(tempFile.Name(), tempFile.Name()+strconv.Itoa(time))
}

func getFileSizeInBytes(tempFile *os.File) int64 {
	fileStat, err := tempFile.Stat()
	if err != nil {
		fmt.Println("Error getting file stats...", err)
	}
	return fileStat.Size()
}
