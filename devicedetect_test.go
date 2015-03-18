package main

import (
	"log"
	"testing"
)

func BenchmarkFindByUserAgent(b *testing.B) {
	var device Device
	fillInUserAgents(devices)
	userAgent := "shaked"
	for n := 0; n < b.N; n++ {
		device = findByUserAgent(userAgent)
	}

	log.Printf("Deviced found: %s", device.Name)
}
