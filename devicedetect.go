package devicedetect

import (
	"net/http"

	"github.com/Shaked/devicedetect/platform"
	"github.com/Shaked/devicedetect/util"
	"github.com/gorilla/context"
)

const (
	UnknownOs      = "UnknownOs"
	UnknownVersion = "UnknownVersion"
)

type DeviceDetect struct {
	r *http.Request
}

func init() {
	injectUserAgents()
}

func NewDeviceDetect(r *http.Request) *DeviceDetect {
	return &DeviceDetect{r: r}
}

func (d *DeviceDetect) FindByUserAgent(userAgent string) platform.Device {
	key := util.UserAgentToKey(userAgent)
	f, ok := compiledUserAgents[key]
	if ok {
		device := f()
		return device
	}
	return nil
}

func (d *DeviceDetect) PlatformType() platform.Device {
	return d.FindByUserAgent(d.r.Header.Get("User-Agent"))
}

// Vars returns the route variables for the current request, if any.
func Platform(r *http.Request) platform.DeviceType {
	if rv := context.Get(r, "Platform"); rv != nil {
		device := rv.(platform.Device)
		return device.Which()
	}
	return platform.UNKNOWN
}

type PlatformHandler interface {
	Mobile(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Tablet(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Desktop(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Tv(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Watch(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
	Bot(w http.ResponseWriter, r *http.Request, d *DeviceDetect)
}

func Handler(h PlatformHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := NewDeviceDetect(r)

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

func HandlerMux(s *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := NewDeviceDetect(r)
		context.Set(r, "Platform", m.PlatformType())
		s.ServeHTTP(w, r)
	})
}
