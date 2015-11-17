package godevicedetect

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"log"
	"net/http"
	"strings"

	"github.com/Shaked/godevicedetect/platform"
	"github.com/Shaked/user-agents/packages/go/gouseragents"
	"github.com/gorilla/context"
)

const (
	UnknownOs      = "UnknownOs"
	UnknownVersion = "UnknownVersion"
)

type PreCompiledHandler struct{}

type Compiled struct {
	hashFunc   string
	userAgents map[string]platform.Device
}

//func() platform.Device
func (p *PreCompiledHandler) compile() *Compiled {
	type BrowserX struct {
		Name    string      `json:name,string,omitempty`
		Version json.Number `json:version,Number,omitempty`
	}
	type EngineX struct {
		Name    string      `json:name,string,omitempty`
		Version json.Number `json:version,Number,omitempty`
	}
	type PlatformX struct {
		Name    string      `json:name,string,omitempty`
		Version json.Number `json:version,Number,omitempty`
		Build   string      `json:build,string,omitempty`
		Model   string      `json:model,string,omitempty`
	}
	type X struct {
		Brand    string      `json:brand,string,omitempty`
		Type     string      `json:type,string,omitempty`
		Platform *PlatformX  `json:platform,omitempty`
		Engine   *EngineX    `json:engine,omitempty`
		Browsers []*BrowserX `json:browsers,omitempty`
		Browser  *BrowserX   `json:browser,omitempty`
	}
	v := struct {
		Meta       map[string]string `json:meta,omitempty`
		UserAgents map[string]X      `json:userAgents,omitempty`
	}{}

	err := json.Unmarshal([]byte(gouseragents.CompiledUserAgents), &v)
	if nil != err {
		panic(err)
	}

	c := &Compiled{}

	hash := v.Meta["hash"]
	c.hashFunc = hash

	userAgents := map[string]platform.Device{}
	var device platform.Device
	for key, deviceInfo := range v.UserAgents {
		deviceName := deviceInfo.Brand
		switch deviceInfo.Type {
		case "desktop":
			device = platform.NewDesktop(deviceName)
			if nil != deviceInfo.Platform {
				device.SetPlatform(platform.NewPlatform(
					deviceInfo.Platform.Name,
					string(deviceInfo.Platform.Version),
				))
			}
			break
		case "tv":
			device = platform.NewTv(deviceName)
			break
		case "tablet":
			device = platform.NewTablet(deviceName)
			break
		case "app":
		case "mobile":
			device = platform.NewMobile(deviceName)
			if nil != deviceInfo.Platform {
				a := platform.NewPlatform(
					deviceInfo.Platform.Name,
					string(deviceInfo.Platform.Version),
				)
				a.SetBuild(string(deviceInfo.Platform.Build))
				a.SetModel(string(deviceInfo.Platform.Model))
				device.SetPlatform(a)
			}
		case "bot":
			device = platform.NewBot(deviceName)
			break
		case "glass":
			device = platform.NewGlass(deviceName)
			break
		default:
			panic("WTF :" + deviceInfo.Type)
		}
		userAgents[key] = device
	}
	c.userAgents = userAgents

	return c
}

type DeviceDetect struct {
	r    *http.Request
	meta *Compiled
}

func NewDeviceDetect(r *http.Request, p *PreCompiledHandler) *DeviceDetect {
	meta := p.compile()
	return &DeviceDetect{r: r, meta: meta}
}

func (d *DeviceDetect) FindByUserAgent(userAgent string) platform.Device {
	key := UserAgentToKey(strings.ToLower(userAgent))
	f, ok := d.meta.userAgents[fmt.Sprint(key)]
	if ok {
		return f
	}
	log.Println("Unknown")
	return platform.NewUnknown(userAgent)
}

func (d *DeviceDetect) PlatformType() platform.Device {
	return d.FindByUserAgent(d.r.Header.Get("User-Agent"))
}

// Vars returns the route variables for the current request, if any.
func Platform(r *http.Request) platform.Device {
	rv := context.Get(r, "Platform")
	device := rv.(platform.Device)
	return device
}

type PlatformHandler interface {
	Mobile(w http.ResponseWriter, r *http.Request, d *platform.DeviceMobile)
	Tablet(w http.ResponseWriter, r *http.Request, d *platform.DeviceTablet)
	Desktop(w http.ResponseWriter, r *http.Request, d *platform.DeviceDesktop)
	Tv(w http.ResponseWriter, r *http.Request, d *platform.DeviceTv)
	Watch(w http.ResponseWriter, r *http.Request, d *platform.DeviceWatch)
	Bot(w http.ResponseWriter, r *http.Request, d *platform.DeviceBot)
	Glass(w http.ResponseWriter, r *http.Request, d *platform.DeviceGlass)
	Unknown(w http.ResponseWriter, r *http.Request, d *platform.DeviceUnknown)
}

func Handler(h PlatformHandler, p *PreCompiledHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := NewDeviceDetect(r, p)
		d := m.FindByUserAgent(r.Header.Get("User-Agent"))
		switch t := m.PlatformType().(type) {
		case *platform.DeviceTablet:
			h.Tablet(w, r, d.(*platform.DeviceTablet))
			break
		case *platform.DeviceMobile:
			h.Mobile(w, r, d.(*platform.DeviceMobile))
			break
		case *platform.DeviceTv:
			h.Tv(w, r, d.(*platform.DeviceTv))
			break
		case *platform.DeviceWatch:
			h.Watch(w, r, d.(*platform.DeviceWatch))
			break
		case *platform.DeviceDesktop:
			h.Desktop(w, r, d.(*platform.DeviceDesktop))
			break
		case *platform.DeviceBot:
			h.Bot(w, r, d.(*platform.DeviceBot))
			break
		case *platform.DeviceGlass:
			h.Glass(w, r, d.(*platform.DeviceGlass))
			break
		case *platform.DeviceUnknown:
			h.Unknown(w, r, d.(*platform.DeviceUnknown))
			break
		default:
			panic(fmt.Sprintf("Type %v doesn't exist", t))
		}

	})
}

func HandlerMux(s *http.ServeMux, p *PreCompiledHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := NewDeviceDetect(r, p)
		context.Set(r, "Platform", m.PlatformType())
		s.ServeHTTP(w, r)
	})
}

func UserAgentToKey(userAgent string) uint32 {
	return crc32.ChecksumIEEE([]byte(userAgent))
}
