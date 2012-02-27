package main

// This file implements fuzzy tagging

import (
	"strings"
)

type Tag struct {
	fuzzy    string
	children []*Tag
	file     string
}

func NewTag(tag string) *Tag {
	return &Tag{Fuzzy(tag), []*Tag{}, ""}
}

// Get a child by fuzzy tag matching.
// If the child does not exist yet, it is added.
func (this *Tag) Child(tag string) *Tag {
	fuzzyTag := Fuzzy(tag)
	for _, c := range this.children {
		if c.fuzzy == fuzzyTag {
			return c
		}
	}
	c := NewTag(tag)
	this.children = append(this.children, c)
	return c
}

func (this *Tag) String() string {
	return this.Print(0)
}

func (this *Tag) Print(indent int) string {
	str := spaces(indent) + this.fuzzy + "\n"
	for _, c := range this.children {
		str += c.Print(indent + 1)
	}
	return str
}

func spaces(howmany int) string {
	return "       "[:howmany]
}

func Fuzzy(file string) string {
	file = strings.ToLower(file)
	return file
}
