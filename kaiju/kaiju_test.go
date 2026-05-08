package kaiju

import "testing"

func TestSummary(t *testing.T) {
	if Summary() == "" {
		t.Fatal("summary should not be empty")
	}
}
