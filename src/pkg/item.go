package clip

import (
	"sort"
	"bytes"
)

type Item struct {
	tag string
	file string
}

type ItemArray []Item


func (arr *ItemArray) Add(item Item){
	*arr = append(*arr, item)
	sort.Sort(arr)
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
