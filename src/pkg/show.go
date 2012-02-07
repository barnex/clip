package clip

// This file implements the "show" command

import ()

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
	}
	ok = true
	return
}
