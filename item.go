package main

import (
	"path"
)

// Item in the library can be a track
// or a collection of other Items.
// Collections represent typically genres, artists, albums, etc.
type Item struct {
	tag      string
	file     string
	children []*Item
}

func NewFile(file string) *Item {
	return &Item{tag: Fuzzy(path.Base(file)), file: file}
}

func (item *Item) String() string {
	return item.tag + ":\t" + item.file
}
