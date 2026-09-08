package main

import (
	"encoding/json"
	"io"
)

type League []Player

func NewLeague(r io.Reader) (League, error) {
	var players League

	err := json.NewDecoder(r).Decode(&players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (l League) Find(name string) *Player {
	for i := range l {
		if l[i].Name == name {
			return &l[i]
		}
	}

	return nil
}
