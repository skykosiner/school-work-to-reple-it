package utils

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func AskForInput(message string) string {
	var stringToReturn string

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s \n", message)
		text, _ := reader.ReadString('\n')

		valid := validateInput(text)

		stringToReturn = text

		if valid {
			break
		}
	}
	return stringToReturn
}

func validateInput(input string) bool {
	// Remove the weird 10 at the end of the text object smh
	input = input[:len(input)-1]

	// Make sure that there is no weird ascii charters
	for i := 0; i < len(input); i++ {
		byteValue := input[i]
		intValue := int(byteValue)

		if intValue < 32 || intValue > 127 {
			return false
		}
	}

	return true
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func GetFilesFromPath(path string) []fs.FileInfo {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}

	defer file.Close()

	fileList, _ := file.Readdir(0)

	return fileList
}
