package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shaked/godevicedetect"
	"github.com/Shaked/godevicedetect/platform"
)

type Handler struct{}

func (h *Handler) Mobile(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceMobile,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}
func (h *Handler) Tablet(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceTablet,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}
func (h *Handler) Desktop(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceDesktop,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}
func (h *Handler) Bot(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceBot,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}
func (h *Handler) Glass(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceGlass,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}
func (h *Handler) Watch(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceWatch,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}
func (h *Handler) Tv(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceTv,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}
func (h *Handler) Unknown(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceUnknown,
) {
	p := d.Platform()
	if nil != p {
		fmt.Fprint(w, "Hello, this is: ", d.Type(), d.Platform().Name(), d.Platform().Version())
	} else {
		fmt.Fprint(w, "Hello, this is: ", d.Type())
	}

}

func main() {
	log.Println("Starting local server http://localhost:8000/ (cmd+click to open from terminal)")
	h := &Handler{}
	mux := http.NewServeMux()
	mux.Handle("/", godevicedetect.Handler(h, nil))
	http.ListenAndServe(":8000", mux)
}
