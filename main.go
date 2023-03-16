package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("/Users/makiuchijunki/Dropbox/macOSIFTTTControl/macaction.txt")
	if err != nil {
		fmt.Println("File read error：", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//Read the file line by line and get the last line
	var lastLine string
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	//Extract the string after "|" from the last line
	index := strings.Index(lastLine, "|")
	if index != -1 {
		fmt.Println(lastLine[index+1:])
	}
}
