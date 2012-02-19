package main

// This file implements the Library data structure.

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Stores a music Library
type Lib struct {
	tracks []*Track
}

// Constructs a new Library
func NewLib() *Lib {
	lib := new(Lib)
	lib.init()
	return lib
}

// Initializes the library
func (lib *Lib) init() {
	lib.tracks = []*Track{}
}

// Recursively import directory or file into library.
func (lib *Lib) Import(arg string) {
	// rm trailing slash
	if strings.HasSuffix(arg, "/") {
		arg = arg[:len(arg)-1]
	}

	info, err := os.Stat(arg)
	Check(err) // TODO: dontcrash

	if info.IsDir() {
		dir, err := os.OpenFile(arg, os.O_RDONLY, 0777)
		Check(err)
		files, err2 := dir.Readdirnames(-1)
		Check(err2)
		for _, f := range files {
			lib.Import(arg + "/" + f)
		}
		return
	}

	if !info.IsDir() {
		lib.tracks = append(lib.tracks, NewTrack(arg))
		return
	}
}

// Print the entire library recursively
func (lib *Lib) WriteTo(out io.Writer) (n int, err error) {
	for _, track := range lib.tracks {
		N, ERR := fmt.Fprintln(out, track)
		if ERR != nil {
			err = ERR
		}
		n += N
	}
	return
}

func (lib *Lib) String() string {
	buf := bytes.NewBuffer([]byte{})
	lib.WriteTo(buf)
	return string(buf.Bytes())
}

// Find items based on tag
//func (lib *Lib) Find(tag string) (items []*Item) {
//	tag = Fuzzy(tag)
//	tracks = []*Item{}
//	for _, track := range lib.tracks {
//		if track.tag == tag {
//			items = append(items, item)
//			Debug("find", tag, ":", item.file)
//		}
//	}
//	return
//}
