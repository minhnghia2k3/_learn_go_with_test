package main

import (
	"errors"
)

var ErrNotFoundWord = errors.New("word not found")
var ErrWordExists = errors.New("word existed")
var ErrWordNotExist = errors.New("word not exist, cannot modify")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) exists(word string) (bool, error) {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFoundWord):
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return "", ErrNotFoundWord
	}

	return val, nil
}

func (d Dictionary) Add(word, value string) error {
	found, err := d.exists(word)
	if err != nil {
		return err
	}

	if found {
		return ErrWordExists
	}

	d[word] = value
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	found, err := d.exists(word)
	if err != nil {
		return err
	}

	if !found {
		return ErrWordNotExist
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	found, err := d.exists(word)
	if err != nil {
		return err
	}
	
	if !found {
		return ErrWordNotExist
	}

	delete(d, word)
	return nil
}
