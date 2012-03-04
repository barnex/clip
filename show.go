package main

// This file implements the "show" command

import (
	"fmt"
)

func init() {
	help["show"] = `Show various types of objects`
}

func (api API) Show(args []string) (resp, err string) {
	for _, arg := range args {
		str, ok := api.player.show(arg)
		if !ok {
			err += "show: not found: " + arg + "\n"
			err += "options: library tree tags <tag>\n"
		}
		resp += str + "\n"
	}
	return
}

func (player *Player) show(str string) (resp string, ok bool) {
	switch str {
	default:
		tag := player.FindTag(str)
		if tag != nil {
			resp = tag.Print(0)
		} else {
			ok = false
			return
		}
	case "library":
		resp = player.Lib.String()
	case "tree":
		resp = player.artists.Print(0)
	case "tags":
		for _, tag := range player.tags {
			resp += fmt.Sprintln(tag)
		}
	}
	ok = true
	return
}
