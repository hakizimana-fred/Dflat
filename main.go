package main

import (
	"Dflat/constants"
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("Getting started")
	content, err := Reader()

	if err != nil {
		slog.Error("Something went wrong while reading file", err)
		return
	}
	constants.PL(content)
}
