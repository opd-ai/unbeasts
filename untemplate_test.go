package untemplate

import "testing"

func TestSummary(t *testing.T) {
	if Summary() != "untemplate template 0.1.0" {
		t.Fatalf("unexpected summary: %q", Summary())
	}
}
