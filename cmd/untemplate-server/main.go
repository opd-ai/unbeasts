package main

import (
	"fmt"
	"net/http"

	"github.com/opd-ai/untemplate"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, untemplate.Summary())
}

func main() {
	http.HandleFunc("/healthz", handler)
	_ = http.ListenAndServe(":8080", nil)
}
