package clip

import (
	"sort"
	"bytes"
)

type Item struct {
	name string
	file *Node
}

type ItemArray []Item

func (arr ItemArray) Lookup(item string) *Node{
	Debug("lookup", item, "in", arr)
//	for i:=range arr{
//		fmt.Println(arr[i])
//	}
	return  library.fs
}

func (arr *ItemArray) Add(item string, file *Node) {
	*arr = append(*arr, Item{item, file})
	sort.Sort(arr)
}

func (arr ItemArray) Len() int {
	return len(arr)
}

func (arr ItemArray) Less(i, j int) bool {
	return Less(arr[i].name, arr[j].name)
}

func Less(i, j string) bool {
	return bytes.Compare([]byte(i), []byte(j)) == -1
}

func (arr ItemArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
