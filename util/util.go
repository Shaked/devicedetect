package util

import "hash/crc32"

func UserAgentToKey(userAgent string) uint32 {
	return crc32.ChecksumIEEE([]byte(userAgent))
}
