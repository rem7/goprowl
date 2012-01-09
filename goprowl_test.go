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
