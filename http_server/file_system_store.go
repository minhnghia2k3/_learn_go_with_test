package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type ReadWriteSeekTruncate interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPLayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func NewFileSystemPlayerStoreFromFile(path string) (store *FileSystemPlayerStore, cleanup func(), error error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file error: %v", err)
	}

	store, err = NewFileSystemPLayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("new fs store error: %v", err)
	}

	cleanup = func() {
		db.Close()
	}

	return store, cleanup, nil
}

func (s *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(s.league, func(i, j int) bool {
		return s.league[i].Wins > s.league[j].Wins
	})
	return s.league
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := s.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	player := s.league.Find(name)

	// init for new user
	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{name, 1})
	}

	s.database.Encode(s.league)
}

func (s *FileSystemPlayerStore) UpdateScore(name string, score int) {}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}
