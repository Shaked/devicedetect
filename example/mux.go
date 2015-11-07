package main

import (
	"fmt"
	"net/http"

	"github.com/Shaked/godevicedetect"
)

func main() {
	h := &StubMuxHandler{}
	m := http.NewServeMux()
	m.Handle("/", h)
	p := &godevicedetect.PreCompiledHandler{}
	http.ListenAndServe(":8000", godevicedetect.HandlerMux(m, p))
}

type StubMuxHandler struct {
	printResult string
}

func (h *StubMuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	device := godevicedetect.Platform(r)
	fmt.Fprintf(w, "Platform-%d-%s: %s", device.Type(), device.Name(), h.printResult)
}
