package file

import (
	"io"
	"os"
	"path/filepath"
)

const absolutePath string = "./internal/file/words.txt"

func ParseAsSlice() []string {
	filePath, err := filepath.Abs(absolutePath)
	checkError(err)

	file, err := os.Open(filePath)
	checkError(err)

	return parseFile(file)
}

func parseFile(file *os.File) []string {
	fileAsSlice := []string{}

	for {
		readBuffer := make([]byte, 5)
		_, err := file.Read(readBuffer)

		if err == io.EOF {
			break
		}
		fileAsSlice = append(fileAsSlice, string(readBuffer))
	}
	return fileAsSlice
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
