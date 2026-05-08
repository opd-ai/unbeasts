package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	handler(rr, httptest.NewRequest("GET", "/healthz", nil))
	if !strings.Contains(rr.Body.String(), "untemplate template") {
		t.Fatalf("unexpected body: %q", rr.Body.String())
	}
}
