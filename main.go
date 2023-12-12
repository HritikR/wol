package main

import (
	"flag"
	"fmt"
	"net"
)

const wolBanner = `
██     ██  ██████  ██      
██     ██ ██    ██ ██      
██  █  ██ ██    ██ ██      
██ ███ ██ ██    ██ ██      
 ███ ███   ██████  ███████ 
                                                      
`

func main() {
	// Display the Wake-on-LAN banner
	fmt.Print(wolBanner)

	// Parse command-line flags
	var targetMAC string
	flag.StringVar(&targetMAC, "m", "", "MAC address of the target device (required)")
	flag.Parse()

	// Validate required flags
	if targetMAC == "" {
		fmt.Println("Error: MAC address is required")
		return
	}

	// Parse MAC address
	hwAddr, err := net.ParseMAC(targetMAC)
	if err != nil {
		fmt.Println("Error parsing MAC address:", err)
		return
	}

	sendMagicPacket(hwAddr)
}