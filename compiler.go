package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

var file = ""

var hashingCollisionErr = errors.New("Hashing collision, try different hashing algorithm")
var unknownBrowser = "Unknown Browser"

type Device struct {
	browser  string
	os       string
	version  string
	platform string
	name     string
}

func NewDevice(browser, os, version, platform, name string) *Device {
	return &Device{browser, os, version, platform, name}
}

func main() {
	var mergedUserAgents map[string]string
	if "" != file {
		newUserAgents := loadNewUserAgentsFromFile(file)
		mergedUserAgents = mergeDevices(userAgents, newUserAgents)
	} else {
		mergedUserAgents = userAgents
	}
	devicesWithData := splitDevicesData(mergedUserAgents)
	hashedDevices := generateHashedDevicesList(devicesWithData)
	log.Println(hashedDevices)
}

// The idea is to compare with the user agent sources to make sure what's new and what not
// Once the list is so big, its really hard to follow updates
func loadNewUserAgentsFromFile(file string) map[string]string {
	newUserAgents := map[string]string{}
	return newUserAgents
}

func mergeDevices(userAgents map[string]string, newUserAgents map[string]string) map[string]string {
	mergedUserAgents := map[string]string{}
	for userAgent, deviceInformation := range userAgents {
		mergedUserAgents[userAgent] = deviceInformation
	}
	return mergedUserAgents
}

func splitDevicesData(userAgents map[string]string) map[string]*Device {
	devicesWithData := map[string]*Device{}
	for userAgent, name := range userAgents {
		name = strings.Trim(name, " ")
		device := deviceFrom(name, userAgent)
		devicesWithData[userAgent] = device
	}
	return devicesWithData
}

func deviceFrom(name, userAgent string) *Device {
	device := checkInHardcodedCase(name, userAgent)
	if nil != device {
		return device
	}

	updatedName, version := findVersion(name)
	//regex to look:
	browser := findBrowser(updatedName, userAgent)
	os := findOs(userAgent)
	platform := findPlatform(userAgent)
	//	find browser:
	//		 name in user agent
	//	find os:
	//		search known os in user agent
	//	find device:
	//		run against device list (taken from mobile detect?)
	//	params:
	//		case insensitive
	//split name to know which browser & version
	//	only if the split returns int\float then it use as version
	//	otherwise use as name
	//
	if "" != browser || "" != os || "" != version || "" != platform || "" != updatedName {
		//log.Printf("Could not get device from name: %s,\n user agent: %s", updatedName, userAgent)
		return nil
	}
	device = NewDevice(browser, os, version, platform, updatedName)
	log.Printf("UserAgent: %s\n device: %+v", userAgent, device)
	return device
}

func checkInHardcodedCase(name, userAgent string) *Device {
	return nil
}

func findPlatform(userAgent string) string {
	platformList := []string{
		"android",
		"blackberry",
		"palm",
		"symbian",
		"windowsmobile",
		"windowsphone",
		"meego",
		"maemo",
		"java",
		"web",
		"bada",
		"brew",
		"iphone",
		"ipad",
		"tablet",
		"mobile",
	}
	for _, platformName := range platformList {
		platformNameStarts := strings.Index(userAgent, platformName)
		if -1 != platformNameStarts {
			platformNameEnds := strings.Index(userAgent[platformNameStarts:], ";")
			if -1 != platformNameEnds {
				platformNameEnds += platformNameStarts
				return userAgent[platformNameStarts:platformNameEnds]
			}
		}
	}

	return ""
}

func findOs(userAgent string) string {
	osList := []string{
		"androidos",
		"blackberryos",
		"palmos",
		"symbianos",
		"windowsmobileos",
		"windowsphoneos",
		"ios",
		"meegoos",
		"maemoos",
		"javaos",
		"webos",
		"badaos",
		"brewos",
		"windows",
		"linux",
		"mac",
	}
	for _, osName := range osList {
		osNameStarts := strings.Index(userAgent, osName)
		if -1 != osNameStarts {
			osNameEnds := strings.Index(userAgent[osNameStarts:], ";")
			if -1 != osNameEnds {
				osNameEnds += osNameStarts
				return userAgent[osNameStarts:osNameEnds]
			}
		}
	}

	return ""

}
func findVersion(name string) (string, string) {
	version := ""
	possibleVersions := strings.Split(name, " ")
	updatedName := name
	for _, pv := range possibleVersions[1:] {
		_, err := strconv.Atoi(pv)
		if nil != err {
			updatedName += pv
			continue
		}

		version = pv
		break
	}
	return updatedName, version
}

//TODO: think about returning also int as how sure the result is
//		e.g first condition is more correct than the second one
//			 and the last is definitely not
func findBrowser(name, userAgent string) string {
	if strings.Contains(strings.ToLower(userAgent), strings.ToLower(name)) {
		return name
	}

	nameWithoutSpaces := strings.Replace(name, " ", "", -1)
	if strings.Contains(strings.ToLower(userAgent), strings.ToLower(nameWithoutSpaces)) {
		return nameWithoutSpaces
	}

	possibleBrowserNames := strings.Split(name, " ")
	for _, pb := range possibleBrowserNames {
		if strings.Contains(strings.ToLower(userAgent), strings.ToLower(pb)) {
			return pb
		}
	}

	return ""
}

func generateHashedDevicesList(devices map[string]*Device) map[uint32]*Device {
	if nil != verifyHashingCollision(devices) {
		panic("hash collision, consider using a different hashing algorithm.")
	}

	return map[uint32]*Device{}
}

func verifyHashingCollision(devices map[string]*Device) error {
	allKeys := map[string]bool{}
	for key := range devices {
		ok := allKeys[key]
		if ok {
			return hashingCollisionErr

		}
		allKeys[key] = true
	}
	return nil
}
