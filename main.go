package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("/Users/makiuchijunki/Dropbox/macOSIFTTTControl/macaction.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Read the file line by line and get the last line
	//var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Extract the string after "|" from the last line
		lastLine := strings.Split(scanner.Text(), "|")
		if len(lastLine) > 1 {
			fmt.Println(lastLine[1])
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
