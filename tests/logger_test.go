package dotenvtest

import (
	"testing"

	"github.com/tonievictor/dotenv"
)

func TestLogger(t *testing.T) {
	t.Run("Test Invalid logger", func(t *testing.T) {
		inValidLoger := "err"

		_, err := dotenv.Config(filenames["valid"],inValidLoger)
		if err == nil {
			t.Error("Expected an err")
		}
	})

	t.Run("Test nil logger", func(t *testing.T) {
		_, err := dotenv.Config(filenames["valid"], nil)
		if err != nil {
			t.Error("Expected an err")
		}
	})

	t.Run("Test nil logger (1 param)", func(t *testing.T) {
		_, err := dotenv.Config(nil)
		if err != nil {
			t.Error("Expected an err")
		}
	})
}
