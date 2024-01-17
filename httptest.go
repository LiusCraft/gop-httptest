package httptest

import (
    "net/http"
	"strings"
    // "net/http/httptest"
    // "testing"
)

type Test struct {
	host string
	port string
	url string
	result *http.Response
}

func (t *Test) Get(url string) {
	t.result, _ = http.Get(url)
}

func (t *Test) Post(url string) {
	t.result, _ = http.Post(url, "application/json", strings.NewReader("name=cjb"))
}