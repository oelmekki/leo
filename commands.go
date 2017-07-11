package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/jwaldrip/odin/cli"
	"github.com/oelmekki/leo/setup"
	"github.com/oelmekki/leo/docker"
	"github.com/oelmekki/leo/letsencrypt"
	"github.com/oelmekki/leo/env"
)

func setupCmd( c cli.Command ) {
	err := setup.Run()

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}

func updateCmd( c cli.Command ) {
	err := setup.Install()

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}

func createCmd( c cli.Command ) {
	err := setup.CreateApp( c.Param( "appName" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func startCmd( c cli.Command ) {
	err := docker.Start( c.Param( "appName" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func stopCmd( c cli.Command ) {
	err := docker.Stop( c.Param( "appName" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func runCmd( c cli.Command ) {
	err := docker.Run( c.Param( "appName" ).String(), c.Param( "service" ).String(), strings.Join( c.Args().Strings(), " " ) )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func rotateCmd( c cli.Command ) {
	err := docker.Rotate( c.Param( "appName" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func letsencryptCmd( c cli.Command ) {
	err := letsencrypt.GetCertificate( c.Param( "appName" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func envCmd( c cli.Command ) {
	err := env.Show( c.Param( "appName" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func envSetCmd( c cli.Command ) {
	err := env.Set( c.Param( "appName" ).String(), c.Param( "variable" ).String(), c.Param( "value" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func envDelCmd( c cli.Command ) {
	err := env.Del( c.Param( "appName" ).String(), c.Param( "variable" ).String() )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}


func implodeCmd( c cli.Command ) {
	err := setup.Implode()

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
}
