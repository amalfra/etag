package etag

import (
	"fmt"
	"testing"
)

func TestNotWeekEtag(t *testing.T) {
	inputStr := `<html>
	<head>
		<title>test page</title>
	</head>
	<body>
		<p>test page html</p>
	</body>
	</html>`
	expectedOut := fmt.Sprintf("%d-%s", len(inputStr), getHash(inputStr))
	generatedOut := Generate(inputStr, false)
	if generatedOut != expectedOut {
		t.Fatalf("Incorrect Etag generated")
	}
}

func TestWeekEtag(t *testing.T) {
	inputStr := `<html>
	<head>
		<title>test page</title>
	</head>
	<body>
		<p>test page html</p>
	</body>
	</html>`
	expectedOut := fmt.Sprintf("W/%d-%s", len(inputStr), getHash(inputStr))
	generatedOut := Generate(inputStr, true)
	if generatedOut != expectedOut {
		t.Fatalf("Incorrect Etag generated")
	}
}
