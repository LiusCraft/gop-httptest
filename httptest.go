package httptest

import (
    "net/http"
	"net/url"
	// "strings"
    // "net/http/httptest"
    // "testing"
)

type Test struct {
	caseName string
	// host string
	// port string
	gUrl string
	result *http.Response
	body map[string][]string
}

func (t *Test) Get(URL string) {
	t.result, _ = http.Get(URL)
	// println(URL)
	// println(t.result)
}

func (t *Test) Post(URL string, body map[string][]string) {
	payload := url.Values(body)
	t.result, _ = http.PostForm(URL, payload)
}
