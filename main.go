package main

import (
	"fmt"
	"log/slog"
)

func main() {
	content, err := Reader()

	if err != nil {
		slog.Error("Something went wrong while reading file", err)
		return
	}

	maxPacketSize := 1024 // 1GB
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

	for _, packet := range packets {
		fmt.Printf("Packet: seqNum=%d, payload=%s\n", packet.seqNum, string(packet.payload))
	}
}
