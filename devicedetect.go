package devicedetect

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"log"
	"net/http"

	"github.com/Shaked/devicedetect/platform"
	"github.com/gorilla/context"
)

const (
	UnknownOs      = "UnknownOs"
	UnknownVersion = "UnknownVersion"
)

type PreCompiledHandler struct{}

//func() platform.Device
func (p *PreCompiledHandler) compile() map[string]interface{} {
	v := map[string]interface{}{}
	log.Println(compiledUserAgents)
	err := json.Unmarshal([]byte(compiledUserAgents), &v)
	if nil != err {
		panic(err)
	}
	log.Println(v)
	return v
}

type DeviceDetect struct {
	r    *http.Request
	meta map[string]interface{}
}

func NewDeviceDetect(r *http.Request, p *PreCompiledHandler) *DeviceDetect {
	meta := p.compile()
	return &DeviceDetect{r: r, meta: meta}
}

func (d *DeviceDetect) FindByUserAgent(userAgent string) platform.Device {
	key := UserAgentToKey(userAgent)
	f, ok := d.meta[fmt.Sprint(key)]
	if ok {
		// device := f()
		device := f
		log.Println("device", device)
		return nil
	}
	return nil
}

func (d *DeviceDetect) PlatformType() platform.Device {
	return d.FindByUserAgent(d.r.Header.Get("User-Agent"))
}

// Vars returns the route variables for the current request, if any.
func Platform(r *http.Request) (platform.DeviceType, string) {
	if rv := context.Get(r, "Platform"); rv != nil {
		device := rv.(platform.Device)
		return device.Which(), device.Name()
	}
	return platform.UNKNOWN, "UNKNOWN"
}

type PlatformHandler interface {
	Mobile(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Tablet(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Desktop(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Tv(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Watch(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Bot(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
}

func Handler(h PlatformHandler, p *PreCompiledHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := NewDeviceDetect(r, p)

		switch m.PlatformType().(type) {
		case *platform.DeviceTablet:
			h.Tablet(w, r, m)
			break
		case *platform.DeviceMobile:
			h.Mobile(w, r, m)
			break
		case *platform.DeviceTv:
			h.Tv(w, r, m)
			break
		case *platform.DeviceWatch:
			h.Watch(w, r, m)
			break
		case *platform.DeviceDesktop:
			h.Desktop(w, r, m)
			break
		case *platform.DeviceBot:
			h.Bot(w, r, m)
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
