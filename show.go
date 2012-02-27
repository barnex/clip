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
		}
		resp += str + "\n"
	}
	return
}

func (player *Player) show(str string) (resp string, ok bool) {
	switch str {
	default:
		return
	case "library":
		resp = player.Lib.String()
	case "tree":
		resp = player.artists.Print(0)
	case "tags":
		resp = fmt.Sprint(player.tags)
	}
	ok = true
	return
}
