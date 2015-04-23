package devicedetect

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shaked/devicedetect/platform"
	"github.com/Shaked/devicedetect/util"
)

var userAgent1 = "ABC-1"
var userAgent2 = "ABC-2"

var compiledUserAgents = map[uint32]func() platform.Device{}

func injectUserAgents() {
	compiledUserAgents[util.UserAgentToKey(userAgent1)] = func() platform.Device {
		return platform.NewDesktop("SomeBrowser-1", "0.6", UnknownOs)
	}
	compiledUserAgents[util.UserAgentToKey(userAgent2)] = func() platform.Device {
		return platform.NewBot("SomeBrowser-2-Bot", UnknownVersion, "Windows NT 6.0")
	}
}

func TestIfHandlerWorksProperly(t *testing.T) {
	h := &StubHandler{printResult: "empty"}
	s := httptest.NewServer(Handler(h))

	uas := map[string]string{
		userAgent1: "Platform-Desktop: empty",
		userAgent2: "Platform-Bot: empty",
	}

	createServer(t, s, s.URL, uas)
}

func TestIfMuxWorsProperly(t *testing.T) {
	m := http.NewServeMux()
	h := &StubMuxHandler{printResult: "empty"}
	m.Handle("/detectdevice", h)
	s := httptest.NewServer(HandlerMux(m))
	u := fmt.Sprintf("%s/detectdevice", s.URL)

	uas := map[string]string{
		userAgent1: "Platform-Desktop: empty",
		userAgent2: "Platform-Bot: empty",
	}

	createServer(t, s, u, uas)
}

func createServer(t *testing.T, s *httptest.Server, u string, uas map[string]string) {
	c := http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if nil != err {
		t.Fatalf("This error should not happen 1: %s", err)
	}

	for ua, expectedBodyString := range uas {
		go func(ua, expectedBodyString string) {
			req.Header.Set("User-Agent", ua)

			res, err := c.Do(req)
			if nil != err {
				t.Fatalf("This error should not happen 2: %s", err)
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if nil != err {
				t.Fatalf("This error should not happen 3: %s", err)
			}

			bodyString := string(body)
			if bodyString != expectedBodyString {
				t.Fatalf("Body is %s instead of %s", bodyString, expectedBodyString)
			}
		}(ua, expectedBodyString)
	}
}
