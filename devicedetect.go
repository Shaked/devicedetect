package main

import (
	"fmt"
	"hash/crc32"
	"log"
)

var devices map[uint32]Device
var deviceToLookup = "shaked"

type Device struct {
	Name    string
	Version string
}

func init() {
	chrome41022270 := Device{"Chrome", "41.0.2227.0"}

	devices = map[uint32]Device{
		crc32.ChecksumIEEE([]byte("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")):                    Device{"Chrome", "41.0.2228.0"},
		crc32.ChecksumIEEE([]byte("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.1 Safari/537.36")): Device{"Chrome", "41.0.2227.1"},
		crc32.ChecksumIEEE([]byte("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36")):                 chrome41022270,
		crc32.ChecksumIEEE([]byte("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")):                    chrome41022270,
	}
}

func main() {
	fillInUserAgents(devices)
	device := findByUserAgent(deviceToLookup)
	log.Println(device)
}

func findByUserAgent(userAgent string) Device {
	return devices[crc32.ChecksumIEEE([]byte(userAgent))]
}

func fillInUserAgents(devices map[uint32]Device) {
	for i := 0; i < 10000000; i++ {
		asIfUserAgent := fmt.Sprintf("this-is-a-user-agent-%d", i)
		asIfVersion := fmt.Sprintf("version%d", i)
		devices[crc32.ChecksumIEEE([]byte(asIfUserAgent))] = Device{asIfUserAgent, asIfVersion}
	}
	devices[crc32.ChecksumIEEE([]byte("shaked"))] = Device{"the shak", "1.2.3.4.5"}
}
