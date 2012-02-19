package main

import(

)


type Track struct{
	file string
}

func NewTrack(file string)*Track{
	track := new(Track)
	track.file = file
	return track
}
