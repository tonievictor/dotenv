package dotenvtest

import (
	"testing"

	"github.com/tonievictor/dotenv"
)

func TestKeyValue(t *testing.T) {
	t.Run("Test invalid key values", func(t *testing.T) {
		filename := dotenv.WithFilename(filenames["invalid-keyval"])

		got, _ := dotenv.Config(filename, dotenv.WithLogger(logger))
		want := map[string]string{"KEY1": "VALUE1"}

		reflectDeepEqual(t, got, want)
	})

	t.Run("Test white space in key value pairs", func(t *testing.T) {
		filename := dotenv.WithFilename(filenames["whitespace"])

		got, _ := dotenv.Config(filename, dotenv.WithLogger(logger))
		want := map[string]string{
			"KEY1": "VALUE1",
			"KEY2": "VALUE2",
		}

		reflectDeepEqual(t, got, want)
	})
}
