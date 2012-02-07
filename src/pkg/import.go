package clip

// This file implements the "import" command

func (api API) Import(args []string) (resp, err string) {
	for _, arg := range args {
		api.player.library.Import(arg)
	}
	return
}
