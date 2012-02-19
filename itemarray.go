package main

import (
	"bytes"
)

type ItemArray []*Item

func (arr *ItemArray) Append(item ...*Item) {
	*arr = append(*arr, item...)
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

func (item Item) String() string {
	return item.tag + ":" + item.file
}
