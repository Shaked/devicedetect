package godevicedetect

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"log"
	"net/http"

	"github.com/Shaked/godevicedetect/platform"
	"github.com/Shaked/user-agents/bin/go/gouseragents"
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
	v := struct {
		Meta       map[string]string                 `json:meta`
		UserAgents map[string]map[string]interface{} `json:userAgents`
	}{}

	log.Println(gouseragents.CompiledUserAgents)
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
		switch deviceInfo["type"].(string) {
		case "desktop":
			device = platform.NewDesktop(deviceInfo["name"].(string))
			break
		case "tablet":
			device = platform.NewTablet(deviceInfo["name"].(string))
			break
		default:
			panic("WTF?")
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
	key := UserAgentToKey(userAgent)
	f, ok := d.meta.userAgents[fmt.Sprint(key)]
	if ok {
		return f
	}
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
	Unknown(w http.ResponseWriter, r *http.Request, d *platform.DeviceUnknown)
}

func Handler(h PlatformHandler, p *PreCompiledHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := NewDeviceDetect(r, p)
		d := m.FindByUserAgent(r.Header.Get("User-Agent"))
		switch m.PlatformType().(type) {
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
