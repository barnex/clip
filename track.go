package main

import(

)


type Track struct{
	file string
	tags [4]string
}

const(
	TAG_TITLE = iota
	TAG_ALBUM
	TAG_ARTIST
	TAG_GENRE
)


func NewTrack(file string)*Track{
	track := new(Track)
	track.file = file
	return track
}

func(track*Track)File()string{
	return track.file
}

func(track*Track)Title()string{
	return track.tags[TAG_TITLE]
}

func(track*Track)Album()string{
	return track.tags[TAG_ALBUM]
}

func(track*Track)Artist()string{
	return track.tags[TAG_ARTIST]
}

func(track*Track)Genre()string{
	return track.tags[TAG_GENRE]
}

func (track*Track)String()string{
	return track.file + "\n\t" + 
		"Title : " + track.Title() + "\n\t" + 
		"Album : " + track.Album() + "\n\t" + 
		"Artist: " + track.Artist() + "\n\t" + 
		"Genre : " + track.Genre() + "\n"
}
