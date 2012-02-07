package clip

// This file implements the "help" command.

import (
	"reflect"
	"unicode"
	"strings"
)

func (player *Player) Help() (resp, err string) {
	p := reflect.ValueOf(player)
	for i := 0; i < p.Type().NumMethod(); i++ {
		m := p.Type().Method(i)
		if isCommand(m){
			resp += strings.ToLower(m.Name) + " "
		}
	}
	return
}

func isCommand(m reflect.Method)bool{
		return unicode.IsUpper(int(m.Name[0])) && m.Type.NumOut() == 2
}
