package env

import (
	"github.com/oelmekki/leo/leo-cli/config"
	"github.com/oelmekki/leo/leo-cli/remote"
)

func Show() ( err error ) {
	cfg, err := config.Read()
	if err != nil { return }

	if err = remote.Run( "leo env " + cfg.AppName ) ; err != nil { return }

	return
}
