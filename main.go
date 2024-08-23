package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

const wolBanner = `
██     ██  ██████  ██      
██     ██ ██    ██ ██      
██  █  ██ ██    ██ ██      
██ ███ ██ ██    ██ ██      
 ███ ███   ██████  ███████ 
                                                      
`

func main() {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// Display the Wake-on-LAN banner
	fmt.Print(wolBanner)

	// Parse command-line flags
	var targetMAC string
	flag.StringVar(&targetMAC, "m", "", "MAC address of the target device (required)")
	flag.Parse()

	// If targetMAC is not supplied, attempt to read from config
	if targetMAC == "" {
		macAddress, err := readConfig()
		if err != nil {
			fmt.Println("Error: MAC address is required and not found in config")
			return
		}
		targetMAC = macAddress
	}

	// Validate that we now have a MAC address
	if targetMAC == "" {
		fmt.Println("Error: MAC address is required")
		flag.Usage()
		os.Exit(1)
	}

	// Parse MAC address
	hwAddr, err := net.ParseMAC(targetMAC)
	if err != nil {
		fmt.Println("Error parsing MAC address:", err)
		return
	}

	if isDeviceOnline(hwAddr) {
		fmt.Println("Device is already online")
		return
	}

	sendMagicPacket(hwAddr)
}
