package dotenvtest

import (
	"testing"

	"github.com/tonievictor/dotenv"
)

func TestFileName(t *testing.T) {
	t.Run("Check for filename (single param)", func(t *testing.T) {
		_, err := dotenv.Config(".env")

		if err != nil {
			t.Error("Single file name shouldn't return an Error")
		}
	})
}
