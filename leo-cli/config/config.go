package config

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Remote struct {
	AppName string
	ServerName string
}

func Read() ( cfg Remote, err error ) {
	content, err := ioutil.ReadFile( "./leo.conf" )
	if err != nil { return cfg, fmt.Errorf( "Can't read config file : %v\nDid you run `leo-cli init` in this repos?", err ) }

	err = json.Unmarshal( content, &cfg )
	if err != nil { return }

	return
}
