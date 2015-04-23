package devicedetect

import "github.com/Shaked/devicedetect/platform"

var compiledUserAgents = map[uint32]func() platform.Device{}

func injectUserAgents() {

	compiledUserAgents[131354474] = func() platform.Device {
		return platform.NewDesktop("ABrowse", "0.6", UnknownOs)
	}
	compiledUserAgents[3750064634] = func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}
	compiledUserAgents[339885753] = func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}
	compiledUserAgents[1891767521] = func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}
	compiledUserAgents[1929521677] = func() platform.Device {
		return platform.NewDesktop("ABrowse", "0.6", UnknownOs)
	}
	compiledUserAgents[1939774630] = func() platform.Device {
		return platform.NewDesktop("ABrowse", "0.4", UnknownOs)
	}
	compiledUserAgents[610906340] = func() platform.Device {
		return platform.NewDesktop("Acoo Browser", "1.98.744", "Windows NT 6.0")
	}
	compiledUserAgents[939591337] = func() platform.Device {
		return platform.NewDesktop("Acoo Browser", "1.98.744", "Windows NT 6.0")
	}
	compiledUserAgents[2443873740] = func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 5.1")
	}
	compiledUserAgents[415810049] = func() platform.Device {
		return platform.NewDesktop("Acoo Browser", UnknownVersion, "Windows NT 6.0")
	}
}
