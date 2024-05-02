package main

import (
	"Dflat/constants"
	"Dflat/utils"
	"bufio"
	"hash/crc32"
	"log"
	"log/slog"
	"os"
	"strings"
)

func Reader() (string, error) {

	file, err := os.Open(constants.FileName)

	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	crc := crc32.NewIEEE()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		content += line + "\n"
		crc.Write([]byte(line))

	}

	fileChecksum := crc.Sum32()
	expectedChecksum, _ := utils.CalculateCheckSum(constants.FileName)

	if fileChecksum != expectedChecksum {
		slog.Error("Encountered file content mismatch, file could be corrupt")

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content, nil
}
