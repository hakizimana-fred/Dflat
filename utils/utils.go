package utils

import (
	"hash/crc32"
	"io"
	"os"
)

func CalculateCheckSum(filename string) (uint32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	crc := crc32.NewIEEE()
	_, err = io.Copy(crc, file)
	if err != nil {
		return 0, err
	}

	return crc.Sum32(), nil
}
