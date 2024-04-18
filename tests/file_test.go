package dotenvtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tonie-ng/go-dotenv"
)

func TestFile(t *testing.T) {
	t.Run("Test valid file", func(t *testing.T) {
		filename := filenames["valid"]

		got, _ := dotenv.Config(filename, logger)
		want := map[string]string{
			"KEY1":           "VALUE1",
			"KEY2":           "VALUE2",
			"DB_CONN_STRING": "postgres://postgres:passw@localhost:5432/dbname?sslmode=disable",
		}

		reflectDeepEqual(t, got, want)
	})

	t.Run("Test nonexistent file", func(t *testing.T) {
		filename := filenames["nonexistent"]

		dotenv.Config(filename, logger)
		err := fmt.Sprintf("open %v: no such file or directory", filename)

		if !strings.Contains(logOutput.String(), err) {
			t.Errorf("Expected %v got %v", err, logOutput.String())
		}
	})

	t.Run("Test empty file", func(t *testing.T) {
		filename := filenames["empty"]

		got, _ := dotenv.Config(filename, logger)

		if len(got) != 0 {
			t.Errorf("Expected an empty map, got %v", got)
		}
	})
}
