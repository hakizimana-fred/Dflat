package main

import (
	"fmt"
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

func (s *Sender) InitializeWindow() {
	s.windowStart = 0
	s.windowEnd = s.windowStart + s.windowSize
	if s.windowEnd > len(s.packets) {
		s.windowEnd = len(s.packets)
	}
}

func (s *Sender) slideWindow() {
	s.windowStart++
	s.windowEnd++
	if s.windowEnd > len(s.packets) {
		s.windowEnd = len(s.packets)
	}
}

func (s *Sender) sendPackets() {
	for i := s.windowStart; i < s.windowEnd; i++ {
		if i >= len(s.packets) {
			break
		}
		SendPacket(s.packets[i])
	}
}

func SendPacket(packet *Packet) {
	fmt.Printf("Sending packet: %s\n", string(packet.payload))
}
