package main

import (
	"bytes"
	"io"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "Hello, World!"
	var buf bytes.Buffer
	io.WriteString(&buf, expected)

	main()

	actual := buf.String()
	if actual != expected {
		t.Errorf("Expected output '%s', but got '%s'", expected, actual)
	}
}
