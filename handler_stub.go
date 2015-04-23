package devicedetect

import (
	"fmt"
	"net/http"
)

type StubHandler struct {
	printResult string
}

func (h *StubHandler) Mobile(w http.ResponseWriter,
	r *http.Request,
	d *DeviceDetect,
) {
	fmt.Fprintf(w, "Platform-%s: %s", "Mobile", h.printResult)
}
func (h *StubHandler) Tablet(w http.ResponseWriter,
	r *http.Request,
	d *DeviceDetect,
) {
	fmt.Fprintf(w, "Platform-%s: %s", "Tablet", h.printResult)
}
func (h *StubHandler) Desktop(w http.ResponseWriter,
	r *http.Request,
	d *DeviceDetect,
) {
	fmt.Fprintf(w, "Platform-%s: %s", "Desktop", h.printResult)
}

func (h *StubHandler) Tv(w http.ResponseWriter,
	r *http.Request,
	d *DeviceDetect,
) {
	fmt.Fprintf(w, "Platform-%s: %s", "Tv", h.printResult)
}
func (h *StubHandler) Watch(w http.ResponseWriter,
	r *http.Request,
	d *DeviceDetect,
) {
	fmt.Fprintf(w, "Platform-%s: %s", "Watch", h.printResult)
}

func (h *StubHandler) Bot(w http.ResponseWriter,
	r *http.Request,
	d *DeviceDetect,
) {
	fmt.Fprintf(w, "Platform-%s: %s", "Bot", h.printResult)
}
