package clip

// This file implements the "import" command

func init(){
	help["import"] = `Import music directory into library`
}

func (api API) Import(args []string) (resp, err string) {
	for _, arg := range args {
		api.player.Import(arg)
	}
	return
}
