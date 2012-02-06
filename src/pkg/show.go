package clip

// This file implements the "show" command

//import (
//	"os"
//	"fmt"
//	"strings"
//)
//
//
//var show map[string]func()string = make(map[string]func()string)
//
//func (player*Player) Show(args []string) (resp string, err os.Error) {
//	// keep map of Showers/Stringers?
//	// also resolve in tags/only if not found?	
//	// show playlist, artists, library, ...
//	object := args[0]
//	f,ok := show[object]
//	if!ok{
//		err = os.NewError("show: not found " + object)
//	}
//	resp = f()
//	return
//}
//
//func showPlaylist() string{
//	return fmt.Sprint(player.playlist)
//}
//
//func showStatus()string{
//	return strings.Replace(fmt.Sprintf("%#v", player), ",", "\n", -1)
//}
