package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Shaked/devicedetect/util"
)

var userAgents = map[string]string{
	`Mozilla/5.0 (compatible; U; ABrowse 0.6; Syllable) AppleWebKit/420+ (KHTML, like Gecko)`: `func() platform.Device {
		return platform.NewDesktop("ABrowse", "0.6", UnknownOs)
	}`,
	`Mozilla/5.0 (compatible; U; ABrowse 0.6;  Syllable) AppleWebKit/420+ (KHTML, like Gecko)`: `func() platform.Device {
		return platform.NewDesktop("ABrowse", "0.6", UnknownOs)
	}`,
	`Mozilla/5.0 (compatible; ABrowse 0.4; Syllable)`: `func() platform.Device {
		return platform.NewDesktop("ABrowse", "0.4", UnknownOs)
	}`,
	`Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; Acoo Browser 1.98.744; .NET CLR 3.5.30729)`: `func() platform.Device {
		return platform.NewDesktop("Acoo Browser", "1.98.744", "Windows NT 6.0")
	}`,
	`Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; Acoo Browser 1.98.744; .NET CLR   3.5.30729)`: `func() platform.Device {
		return platform.NewDesktop("Acoo Browser", "1.98.744", "Windows NT 6.0")
	}`,
	`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0;   Acoo Browser; GTB5; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;   SV1) ; InfoPath.1; .NET CLR 3.5.30729; .NET CLR 3.0.30618)`: `func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}`,
	`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; SV1; Acoo Browser; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; Avant Browser)`: `func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 5.1")
	}`,
	`Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1;   .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)`: `func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}`,
	`Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; GTB5; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; Maxthon; InfoPath.1; .NET CLR 3.5.30729; .NET CLR 3.0.30618)`: `func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}`,
	`Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; GTB5;`: `func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}`,
}

func main() {
	ioutil.WriteFile("./compiled_devices.go", Compile(), 0644)
}

func Compile() []byte {

	code := `
	package devicedetect

	const (
		UnknownOs      = "UnknownOs"
		UnknownVersion = "UnknownVersion"
	)

	var compiledUserAgents = map[uint32]func() platform.Device{}

	func injectUserAgents(){

`
	for ua, fDevice := range userAgents {
		key := util.UserAgentToKey(ua)
		row := fmt.Sprintf("compiledUserAgents[%d]=%s\n", key, fDevice)
		code += row
	}

	code += "}"

	log.Println(code)
	return []byte(code)
}
