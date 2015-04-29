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
var userAgent3 = "ABC-3"
var userAgent4 = "ABC-4"
var userAgent5 = "ABC-5"
var userAgent6 = "ABC-6"

var compiledUserAgents = map[uint32]func() platform.Device{}

type CompiledMock struct{}

func (c *CompiledMock) InjectUserAgents() {
	compiledUserAgents[util.UserAgentToKey(userAgent1)] = func() platform.Device {
		return platform.NewDesktop("SomeBrowser-1", "0.6", UnknownOs)
	}
	compiledUserAgents[util.UserAgentToKey(userAgent2)] = func() platform.Device {
		return platform.NewBot("SomeBrowser-2-Bot", UnknownVersion, "Windows NT 6.0")
	}
	compiledUserAgents[util.UserAgentToKey(userAgent3)] = func() platform.Device {
		return platform.NewMobile("SomeBrowser-3-Mobile", UnknownVersion, "Android 5.1")
	}
	compiledUserAgents[util.UserAgentToKey(userAgent4)] = func() platform.Device {
		return platform.NewTablet("SomeBrowser-4-Tablet", UnknownVersion, "Samsung Galaxy Tab")
	}
	compiledUserAgents[util.UserAgentToKey(userAgent5)] = func() platform.Device {
		return platform.NewTv("SomeBrowser-5-Bot", UnknownVersion, "Apple Tv")
	}
	compiledUserAgents[util.UserAgentToKey(userAgent6)] = func() platform.Device {
		return platform.NewWatch("SomeBrowser-6-Bot", UnknownVersion, "Apple Watch")
	}

}

func TestSearchForUserAgentX(t *testing.T) {
	h := &StubHandler{printResult: "empty"}
	s := httptest.NewServer(Handler(h, nil))
	uas := map[string]string{
		"unknown": "",
	}

	createServer(t, s, s.URL, uas)
}

func TestIfHandlerWorksProperly(t *testing.T) {
	h := &StubHandler{printResult: "empty"}
	s := httptest.NewServer(Handler(h, &CompiledMock{}))

	uas := map[string]string{
		userAgent1: "Platform-Desktop: empty",
		userAgent2: "Platform-Bot: empty",
		userAgent3: "Platform-Mobile: empty",
		userAgent4: "Platform-Tablet: empty",
		userAgent5: "Platform-Tv: empty",
		userAgent6: "Platform-Watch: empty",
	}

	createServer(t, s, s.URL, uas)
}

func TestIfMuxWorsProperly(t *testing.T) {
	m := http.NewServeMux()
	h := &StubMuxHandler{printResult: "empty"}
	m.Handle("/detectdevice", h)
	s := httptest.NewServer(HandlerMux(m, &CompiledMock{}))
	u := fmt.Sprintf("%s/detectdevice", s.URL)

	uas := map[string]string{
		userAgent1:     fmt.Sprintf("Platform-%d: empty", platform.DESKTOP),
		userAgent2:     fmt.Sprintf("Platform-%d: empty", platform.BOT),
		"doesnt-exist": fmt.Sprintf("Platform-%d: empty", platform.UNKNOWN),
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
	}
}
