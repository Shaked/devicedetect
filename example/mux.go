package main

import (
	"fmt"
	"net/http"

	"github.com/Shaked/devicedetect"
)

func main() {
	h := &StubMuxHandler{}
	m := http.NewServeMux()
	m.Handle("/", h)
	p := &devicedetect.PreCompiledHandler{}
	http.ListenAndServe(":8000", devicedetect.HandlerMux(m, p))
}

type StubMuxHandler struct {
	printResult string
}

func (h *StubMuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	device := devicedetect.Platform(r)
	fmt.Fprintf(w, "Platform-%d-%s: %s", device.Type(), device.Name(), h.printResult)
}
