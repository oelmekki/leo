package letsencrypt

import (
	"github.com/oelmekki/leo/leo-cli/config"
	"github.com/oelmekki/leo/leo-cli/remote"
	"github.com/oelmekki/leo/leo-cli/env"
	"fmt"
)

func Run() ( err error ) {
	_, err = env.Get( "LETS_ENCRYPT_EMAIL" )
	if err != nil {
		if err = askForEmail() ; err != nil { return }
	}

	cfg, err := config.Read()
	if err != nil { return }

	if err = remote.Run( "leo letsencrypt " + cfg.AppName ) ; err != nil { return }

	fmt.Println( "SSL ready!" )

	return
}
