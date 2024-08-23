# Wake-on-LAN (WoL) CLI Tool

A simple Wake-on-LAN (WoL) command-line interface (CLI) tool written in Go. WoL allows you to remotely power up devices on your local network.

```
            ██     ██  ██████  ██      
            ██     ██ ██    ██ ██      
            ██  █  ██ ██    ██ ██      
            ██ ███ ██ ██    ██ ██      
            ███ ███   ██████  ███████ 
```

## Build

To build the Wake-on-LAN CLI tool, ensure you have Go installed on your system. Run the following command in your terminal:

```bash
make
```

This will create an executable named `wol`.

## Usage

### 1. Run with MAC Address as Argument

```bash
./wol -m <MAC_ADDRESS>
```

Replace `<MAC_ADDRESS>` with the MAC address of the target device you want to wake up.

### Example

```bash
./wol -m 00:11:22:33:44:55
```

### 2. Run with Configuration File

Add `config.json` file in the root directory with the following format:
```json

{
    "macAddress": "00:11:22:33:44:55"
}

```

## Options

- **-m:** MAC address of the target device (required).
