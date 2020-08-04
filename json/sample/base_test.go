package sample

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

// Package json implements encoding and decoding of JSON as defined in RFC 7159.

func TestCompat(t *testing.T) {
	goodJSON := `{"1": 1}`
	dst := bytes.Buffer{}
	json.Compact(&dst, []byte(goodJSON))
	fmt.Println(dst)
	//{[123 34 49 34 58 49 125] 0 0}
}

func TestIndent(t *testing.T) {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)
}

func TestUnmarshalIndent(t *testing.T) {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	json, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}

func TestValid(t *testing.T) {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))
}

//HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029 characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029 so that the JSON will be safe to embed inside HTML <script> tags. For historical reasons, web browsers don't honor standard HTML escaping within <script> tags, so an alternative JSON encoding must be used.
func TestHTMLEscape(t *testing.T) {
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)
}
