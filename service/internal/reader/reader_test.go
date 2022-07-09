package reader

import (
	"bytes"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("it should return url from file buffer", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.WriteString("__URL__")

		got, err := ReadFile(&buffer)
		if err != nil {
			t.Error("Expected no err, but there is err")
		}

		want := "__URL__"
		if got[0] != want {
			t.Errorf("got %q, want %q", got[0], want)
		}
	})
}
