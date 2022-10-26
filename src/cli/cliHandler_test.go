package cli

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	cmd, content := parseInput("This is a test")
	if cmd != "This" {
		t.FailNow()
	}
	if content != "is a test" {
		t.FailNow()
	}
}

func TestParseInvalidInput(t *testing.T) {
	_, content := parseInput("a")
	if content != "" {
		t.FailNow()
	}
}
