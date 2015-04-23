package devicedetect

import (
	"fmt"
	"net/http"
)

type StubMuxHandler struct {
	printResult string
}

func (h *StubMuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Platform-%s: %s", Platform(r), h.printResult)
}
