package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	file, err := os.Open("/Users/makiuchijunki/Dropbox/macOSIFTTTControl/macaction.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	//Read the file line by line and get the last line
	scanner := bufio.NewScanner(file)
	if scannerErr := scanner.Err(); scannerErr != nil {
		log.Fatal(err)
		return
	}
	for scanner.Scan() {
		//Extract the string after "|" from the last line
		lastLine := strings.Split(scanner.Text(), "|")

		//Argument to mac action
		if lastLine[1] == "sleep" {
			macsleep := exec.Command("sudo", "pmset", "sleepnow")
			macsleepErr := macsleep.Run()
			if macsleepErr != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("Successfully put macOS to sleep.")
		}
	}
}
