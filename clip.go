package main

import (
	"path"
	"unicode"
)

type Clip struct {
	file string
	tags [5]string
}

const (
	TAG_TRACK = iota
	TAG_TITLE
	TAG_ALBUM
	TAG_ARTIST
	TAG_GENRE
)

func NewClip(file string) *Clip {
	clip := new(Clip)
	clip.file = file

	// set num tag
	base := path.Base(file)
	i := 0
	for _, chr := range base {
		if !unicode.IsDigit(chr) {
			break
		}
		i++
	}
	Debug("clip.tags[TAG_TRACK] =",base[:i])
	clip.tags[TAG_TRACK] = base[:i]

	//ext := path.Ext(base)

	return clip
}

func (clip *Clip) File() string {
	return clip.file
}

func (clip *Clip) Track() string {
	return clip.tags[TAG_TRACK]
}

func (clip *Clip) Title() string {
	return clip.tags[TAG_TITLE]
}

func (clip *Clip) Album() string {
	return clip.tags[TAG_ALBUM]
}

func (clip *Clip) Artist() string {
	return clip.tags[TAG_ARTIST]
}

func (clip *Clip) Genre() string {
	return clip.tags[TAG_GENRE]
}

func (clip *Clip) String() string {
	return clip.file + "\n\t" +
		"Track : " + clip.Track() + "\n\t" +
		"Title : " + clip.Title() + "\n\t" +
		"Album : " + clip.Album() + "\n\t" +
		"Artist: " + clip.Artist() + "\n\t" +
		"Genre : " + clip.Genre() + "\n"
}
