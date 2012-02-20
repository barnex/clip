package main

import (
	"path"
	"unicode"
)

// Represents a music clip.
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
	clip.initTags()
	return clip
}

// Rudimentary way to set clip tags based on file name:
//	artist/album/01_title.ogg
// TODO: read I3D tags.
func(clip*Clip)initTags(){
	// if file starts with number,
	// use it as TRACK tag.
	file := clip.file
	base := path.Base(file)
	i := 0
	for _, chr := range base {
		if !unicode.IsDigit(chr) {
			break
		}
		i++
	}
	clip.tags[TAG_TRACK] = base[:i]

	// TITLE tag is filename without extension
	// or leading track number.
	ext := path.Ext(base)
	clip.tags[TAG_TITLE] = base[i:len(base)-len(ext)]

	// ALBUM tag is clip's parent directory
	parent1, _ := path.Split(file)
	clip.tags[TAG_ALBUM] = path.Base(parent1)	

	// ARTIST tag is albums' parent directory
	parent2, _ := path.Split(parent1[:len(parent1)-1])
	clip.tags[TAG_ARTIST] = path.Base(parent2)	
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


