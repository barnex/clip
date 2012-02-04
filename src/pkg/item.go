package clip

import (
	"bytes"
	"path"
)

// Item in the library can be a track
// or a collection of other Items.
// Collections represent typically genres, artists, albums, etc.
type Item struct {
	tag string
	file string
	// []*Item
}

type ItemArray []*Item


func NewFile(file string)*Item{
	return &Item{tag: Fuzzy(path.Base(file)), file:file}
}

func (arr ItemArray) Len() int {
	return len(arr)
}

func (arr ItemArray) Less(i, j int) bool {
	return Less(arr[i].tag, arr[j].tag)
}

func Less(i, j string) bool {
	return bytes.Compare([]byte(i), []byte(j)) == -1
}

func (arr ItemArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func(item Item)String()string{
	return item.tag + ":" + item.file
}
