package poker

import (
	"io"
	"testing"
)

func TestTapteWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("Abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "Abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
