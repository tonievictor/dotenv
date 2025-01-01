package dotenvtest

import (
	"reflect"
	"testing"
	"github.com/tonievictor/dotenv"
)


func TestKeyValue(t *testing.T) {

	t.Run("Test invalid key values", func(t *testing.T) {
		filename := filenames["invalid-keyval"]

		got, _ := dotenv.Config(filename, logger)
		want := map[string]string{"KEY1": "VALUE1"}

		reflectDeepEqual(t, got, want)
	})

	t.Run("Test white space in key value pairs", func(t *testing.T) {
		filename := filenames["whitespace"]

		got, _ := dotenv.Config(filename, logger)
		want := map[string]string{
			"KEY1": "VALUE1",
			"KEY2": "VALUE2",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}
