package main

// The tag tree is used to look up music clips
// using fuzzy matching (so that the name does not
// have to be typed 100% correct on the command line)
// A tag may represent, e.g., an artist, album, clip...
// Some tags, like albums, have child Tags while
// others, like clips, are leaf nodes.
// The leaf nodes point to a file to be played.

import (
	"strings"
)

type Tag struct {
	fuzzy    string // fuzzyfied tag
	children []*Tag // children, if any
	file     string // music file, in case of leaf node
}

func NewTag(tag string) *Tag {
	return &Tag{Fuzzy(tag), []*Tag{}, ""}
}

// Get a child by fuzzy tag matching.
// If the child does not exist yet, it is added.
func (this *Tag) Child(tag string) (child *Tag, ok bool) {
	fuzzyTag := Fuzzy(tag)
	for _, c := range this.children {
		if c.fuzzy == fuzzyTag {
			child = c
			ok = true
			return
		}
	}
	child = NewTag(tag)
	this.children = append(this.children, child)
	return
}

func (this *Tag) String() string {
	return this.fuzzy
}

func (this *Tag) Print(indent int) string {
	str := spaces(indent) + this.fuzzy + ":" + this.file + "\n"
	for _, c := range this.children {
		str += c.Print(indent + 1)
	}
	return str
}

func spaces(howmany int) string {
	return "       "[:howmany]
}

// Fuzzyfy string
func Fuzzy(str string) string {
	str = strings.ToLower(str)
	return str
}
