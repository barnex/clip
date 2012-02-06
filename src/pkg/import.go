package clip

// This file implements the "import" command


func (player *Player) Import(args []string) (resp, err string) {
	if len(args) == 0 {
		err = "nothing specified, nothing imported"
		return
	}
	for _, arg := range args {
		player.library.Import(arg)
	}
	return
}
