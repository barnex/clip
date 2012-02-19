all: clip

clip: *.go
	go tool 6g -o clip.6 *.go
	go tool 6l -o clip clip.6

opt: *.go
	go tool 6g -m -o clip.6 *.go
	go tool 6l -o clip clip.6
