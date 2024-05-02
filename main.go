package main

import (
	"log/slog"
)

func main() {
	_, err := Reader()

	if err != nil {
		slog.Error("Something went wrong while reading file", err)
		return
	}

}
