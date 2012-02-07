package clip

// This file implements the "help" command.

import (
	"reflect"
	"strings"
)

// Store help for commands here
var help map[string]string = make(map[string]string)

func init(){
	help["help"] = `Display this help message`
}

// report all methods on API, using reflection
func (api API) Help() (resp, err string) {
	p := reflect.ValueOf(api)
	for i := 0; i < p.Type().NumMethod(); i++ {
		m := p.Type().Method(i)
		name := strings.ToLower(m.Name)
		resp += "\t" + name + "\t" + help[name] + "\n"
	}
	return
}

