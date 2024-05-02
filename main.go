package main

import (
	"log/slog"
)

func main() {
	content, err := Reader()

	if err != nil {
		slog.Error("Something went wrong while reading file", err)
		return
	}

	maxPacketSize := 1024
	packets := make([]*Packet, 0)

	for i := 0; i < len(content); i += maxPacketSize {
		end := i + maxPacketSize

		if end > len(content) {
			end = len(content)
		}

		packetData := content[i:end]
		packet := &Packet{payload: []byte(packetData)}
		packets = append(packets, packet)

	}

	sender := Sender{
		windowSize:   5,
		packets:      packets,
		windowStart:  0,
		windowEnd:    4,
		ackReceived:  -1,
		sendComplete: false,
	}

	sender.Send()
}
