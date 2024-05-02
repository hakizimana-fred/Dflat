package main

import (
	"fmt"
	"time"
)

type Packet struct {
	seqNum  int
	payload []byte
}

type Sender struct {
	windowSize   int
	packets      []*Packet
	nextPacket   int
	ackReceived  int
	windowStart  int
	windowEnd    int
	sendComplete bool
}

func (s *Sender) Send() {
	for !s.sendComplete {
		for i := s.windowStart; i <= s.windowEnd; i++ {
			if i < len(s.packets) {
				SendPacket(s.packets[i])
			}
		}

		if s.windowEnd >= len(s.packets)-1 {
			s.sendComplete = true
			break
		}

		time.Sleep(1 * time.Second)
		s.slideWindow()

	}

}

func (s *Sender) slideWindow() {
	for s.windowStart <= s.ackReceived {
		s.windowStart++
	}

	s.windowEnd = s.windowStart + s.windowSize - 1

	if s.windowEnd >= len(s.packets) {
		s.windowEnd = len(s.packets) - 1
	}
}

func SendPacket(packet *Packet) {
	fmt.Printf("Sending packet: %s\n", string(packet.payload))
}
