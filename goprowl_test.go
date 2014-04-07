package goprowl

import (
	"strings"
	"testing"
)

func TestErrorParsing(t *testing.T) {
	errRes := `<?xml version="1.0" encoding="UTF-8"?>
<prowl>
<error code="401">Invalid API key(s).</error>
</prowl>`

	expected := "Invalid API key(s)."

	err := decodeError("x", strings.NewReader(errRes))
	if err.Error() != expected {
		t.Fatalf("Expected %s, got %s", expected, err.Error())
	}
}

func TestRegisterKey(t *testing.T) {
	var p Goprowl
	if err := p.RegisterKey("12345"); err == nil {
		t.Fatalf("Register keys not filtering key lenght properly")
	}

	if err := p.RegisterKey("1234512345123451234512345123451234512345"); err != nil {
		t.Fatalf("Register keys not working properly")
	}
}

func TestDelKey(t *testing.T) {
	var p Goprowl

	if err := p.DelKey("12345"); err == nil {
		t.Fatalf("DelKey allows deletion of keys that don't exist")
	}

	err := p.RegisterKey("1234512345123451234512345123451234512345")
	if err != nil {
		t.Fatalf("Register keys not working properly")
	}

	if len(p.apikeys) != 1 {
		t.Fatalf("Register keys not working properly")
	}

	if err := p.DelKey("1234512345123451234512345123451234512345"); err != nil {
		t.Fatalf("DelKey isn't working properly")
	}

	if len(p.apikeys) != 0 {
		t.Fatalf("DelKey isn't working properly")
	}
}
