package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/jwaldrip/odin/cli"
	"github.com/oelmekki/leo/leo-cli/appinit"
	"github.com/oelmekki/leo/leo-cli/deploy"
	"github.com/oelmekki/leo/leo-cli/env"
	"github.com/oelmekki/leo/leo-cli/letsencrypt"
	"github.com/oelmekki/leo/leo-cli/remote"
)

func initCmd( c cli.Command ) {
	if err := appinit.Run( c.Param( "appName" ).String(), c.Param( "server" ).String() ) ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}
}

func deployCmd( c cli.Command ) {
	if err := deploy.Run() ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}
}

func runCmd( c cli.Command ) {
	if err := remote.RunInteractive( c.Param( "service" ).String(), strings.Join( c.Args().Strings(), " " ) ) ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}
}

func envCmd( c cli.Command ) {
	if err := env.Show() ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}
}

func envSetCmd( c cli.Command ) {
	if err := env.Set( c.Param( "variable" ).String(), c.Param( "value" ).String() ) ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}
}

func envDelCmd( c cli.Command ) {
	if err := env.Del( c.Param( "variable" ).String() ) ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}
}

func letsencryptCmd( c cli.Command ) {
	if err := letsencrypt.Run() ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}
}
