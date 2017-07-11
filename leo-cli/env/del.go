package env

import (
	"github.com/oelmekki/leo/leo-cli/config"
	"github.com/oelmekki/leo/leo-cli/remote"
)

func Del( variable string ) ( err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	if err = remote.Run( "leo env:del " + cfg.AppName + " " + variable ) ; err != nil { return }

	return
}
