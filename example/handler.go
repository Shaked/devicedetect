package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shaked/devicedetect"
	"github.com/Shaked/devicedetect/platform"
)

type Handler struct{}

func (h *Handler) Mobile(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceMobile,
) {
	fmt.Fprint(w, "Hello, this is: ", d.Type())
}
func (h *Handler) Tablet(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceTablet,
) {
	fmt.Fprint(w, "Hello, this is: ", d.Type())
}
func (h *Handler) Desktop(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceDesktop,
) {
	fmt.Fprint(w, "Hello, this is: ", d.Type())
}
func (h *Handler) Bot(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceBot,
) {
	fmt.Fprint(w, "Hello, this is: ", d.Type())
}
func (h *Handler) Watch(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceWatch,
) {
	fmt.Fprint(w, "Hello, this is: ", d.Type())
}
func (h *Handler) Tv(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceTv,
) {
	fmt.Fprint(w, "Hello, this is: ", d.Type())
}
func (h *Handler) Unknown(w http.ResponseWriter,
	r *http.Request,
	d *platform.DeviceUnknown,
) {
	fmt.Fprint(w, "Hello, this is: ", d.Type())
}

func main() {
	log.Println("Starting local server http://localhost:8000/ (cmd+click to open from terminal)")
	h := &Handler{}
	mux := http.NewServeMux()
	mux.Handle("/", devicedetect.Handler(h, nil))
	http.ListenAndServe(":8000", mux)
}
