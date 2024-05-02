package main

import (
	"Dflat/constants"
	"bufio"
	"log"
	"os"
)

func Reader() (string, error) {

	file, err := os.Open(constants.FileName)

	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// if err := file.Close(); err != nil {
	// 	log.Fatal("Error closing file:", err)
	// }

	return content, nil

}
