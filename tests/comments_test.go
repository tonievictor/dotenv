package dotenvtest

import (
	"testing"

	"github.com/tonie-ng/go-dotenv"
)

func TestComments(t *testing.T) {
	t.Run("Test comments at the start of a line", func(t *testing.T) {
		filename := filenames["comment"]

		got, _ := dotenv.Config(filename, logger)
		want := map[string]string{
			"KEY1":           "VALUE1",
			"DB_CONN_STRING": "postgres://postgres:passw@localhost:5432/dbname?sslmode=disable",
		}

		reflectDeepEqual(t, got, want)
	})

	t.Run("Test inline comments", func(t *testing.T) {
		filename := filenames["inline-comment"]

		got, _ := dotenv.Config(filename, logger)
		want := map[string]string{
			"KEY1":           "VALUE1",
			"KEY2":           "VALUE2",
			"DB_CONN_STRING": "postgres://postgres:passw@localhost:5432/dbname?sslmode=disable",
		}

		reflectDeepEqual(t, got, want)
	})
}
