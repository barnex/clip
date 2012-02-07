package clip

// This file implements the "help" command.

import (
	"reflect"
	"strings"
)

func (api API) Help() (resp, err string) {
	p := reflect.ValueOf(api)
	for i := 0; i < p.Type().NumMethod(); i++ {
		m := p.Type().Method(i)
		resp += strings.ToLower(m.Name) + " "
	}
	return
}

