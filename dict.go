package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	_, ok := d[word]
	if !ok {
		return "", errors.New("Couldn't find word")
	}
	return d[word], nil
}

func Search(dict map[string]string, key string) string {
	return dict[key]
}
