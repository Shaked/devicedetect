package main

import (
	"hash/crc32"
	"log"
)

var deviceToLookup = "shaked"

func main() {
	//fillInUserAgents(devices)
	device := findByUserAgent(deviceToLookup)
	log.Println(device)
}

func findByUserAgent(userAgent string) Device {
	return devices[crc32.ChecksumIEEE([]byte(userAgent))]
}
