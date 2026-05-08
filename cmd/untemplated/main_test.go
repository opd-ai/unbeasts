package main

import "testing"

func TestRun(t *testing.T) {
	if run() == "" {
		t.Fatal("run() should return template summary")
	}
}
