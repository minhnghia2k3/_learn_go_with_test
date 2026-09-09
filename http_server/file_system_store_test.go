package poker

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPLayerStore(database)
		assertNoError(t, err)
		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPLayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEqual(t, got, want)
	})

	t.Run("league from a reader", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPLayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPLayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEqual(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, removeFile := createTempFile(t, "")
		defer removeFile()

		_, err := NewFileSystemPLayerStore(database)

		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPLayerStore(database)

		assertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	file, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatal("cannot create file", err)
	}

	// make sure cleanup temp file
	cleanup := func() {
		file.Close()
		os.Remove(file.Name())
	}

	if _, err := file.Write([]byte(initialData)); err != nil {
		t.Fatal("write error", err)
	}

	return file, cleanup
}

func assertScoreEqual(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expect no error but got one, %v", err)
	}
}
