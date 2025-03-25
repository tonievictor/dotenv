package dotenvtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tonievictor/dotenv"
)

func TestFile(t *testing.T) {
	t.Run("Test valid file", func(t *testing.T) {
		filename := dotenv.WithFilename(filenames["valid"])

		got, _ := dotenv.Config(filename, dotenv.WithLogger(logger))
		want := map[string]string{
			"KEY1":           "VALUE1",
			"KEY2":           "VALUE2",
			"DB_CONN_STRING": "postgres://postgres:passw@localhost:5432/dbname?sslmode=disable",
		}

		reflectDeepEqual(t, got, want)
	})

	t.Run("Test nonexistent file", func(t *testing.T) {
		filename := dotenv.WithFilename(filenames["nonexistent"])

		dotenv.Config(filename, dotenv.WithLogger(logger))
		err := fmt.Sprintf("open %v: no such file or directory", filename)

		if !strings.Contains(logOutput.String(), err) {
			t.Errorf("Expected %v got %v", err, logOutput.String())
		}
	})

	t.Run("Test empty file", func(t *testing.T) {
		filename := dotenv.WithFilename(filenames["empty"])

		got, _ := dotenv.Config(filename, dotenv.WithLogger(logger))

		if len(got) != 0 {
			t.Errorf("Expected an empty map, got %v", got)
		}
	})
}
