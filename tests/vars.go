package dotenvtest

import (
	"bytes"
	"log"
	"reflect"
	"testing"
)

var logOutput bytes.Buffer
var logger = log.New(&logOutput, "Test logger:", log.LstdFlags)
var filenames = map[string]string{
	"valid": "env/valid.env",
	"nonexistent": "env/nonexistent.env",
	"empty": "env/empty.env",
	"invalid-keyval": "env/invalid-keyval.env",
	"whitespace": "env/whitespace.env",
	"comment": "env/comments.env",
	"inline-comment": "env/inline-comment.env",
}

func reflectDeepEqual(t testing.TB, got, want map[string]string) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected: %v, Got: %v", want, got)
		}
}
