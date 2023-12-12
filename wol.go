package main

import (
	"fmt"
	"net"
)

func sendMagicPacket(hwAddr net.HardwareAddr) {
	// Create magic packet
	magicPacket := generateMagicPacket(hwAddr)

	// Send magic packet to the broadcast address
	broadcastAddr := net.IPv4bcast
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: broadcastAddr, Port: 9})
	if err != nil {
		fmt.Println("Error opening UDP connection:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(magicPacket)
	if err != nil {
		fmt.Println("Error sending magic packet:", err)
		return
	}

	fmt.Println("Magic packet sent successfully to", hwAddr.String())
}

func generateMagicPacket(hwAddr net.HardwareAddr) []byte {
	// Magic packet structure: 6 bytes of 0xFF followed by 16 repetitions of the target MAC address
	packet := make([]byte, 6+16*len(hwAddr))
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}
	for i := 6; i < len(packet); i += len(hwAddr) {
		copy(packet[i:i+len(hwAddr)], hwAddr)
	}

	return packet
}