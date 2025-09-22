package keefies

import (
	"os"
	"testing"

	"pgregory.net/rapid"
)

func TestGetEnv(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		key := rapid.StringMatching("[a-zA-Z_]{1,}[a-zA-Z0-9_]{0,}").Draw(t, "key")

		value := rapid.OneOf(rapid.StringMatching("[a-zA-Z_]{1,}[a-zA-Z0-9_]{0,}"), rapid.StringN(0, 0, 0)).Draw(t, "value")

		shouldSet := value != ""

		if shouldSet {
			t.Logf("Setting %s to %s", key, value)
			err := os.Setenv(key, value)

			if err != nil {
				panic(err)
			}
		}

		result, ok := GetEnv(key)

		if shouldSet && !ok {
			t.Fatalf("GetEnv isn't getting a set env")
		}

		if !shouldSet && ok {
			t.Fatalf("GetEnv is getting the incorrect env")
		}

		if shouldSet && ok && result != value {
			t.Fatalf("GetEnv is returning an incorrect value")
		}
	})
}
