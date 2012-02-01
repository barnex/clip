package clip

// This file implements fuzzy tagging

import (
	"path"
	"strings"
)

func Tag(file string)string{
	file = path.Base(file)
	file = strings.ToLower(file)
	return file
}
