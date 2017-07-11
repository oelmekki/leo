package main

import (
	"os"
	"fmt"
	"github.com/jwaldrip/odin/cli"
	"github.com/oelmekki/leo/validator"
)

func defaultUsage( c cli.Command ) {
	c.Usage()
}
var app = cli.New( "1.0.0", "leo manager", defaultUsage )

/*
 * Usual way to setup odin's accepted commands
 */
func init() {
	app.DefineSubCommand( "setup", "Install leo on this server", setupCmd )
	app.DefineSubCommand( "update", "Install only the leo binary, useful after update", updateCmd )

	subcmd := app.DefineSubCommand( "create", "Create a new app (-h for more)", createCmd )
	subcmd.DefineParams( "appName" )

	subcmd = app.DefineSubCommand( "start", "Run <appName>, if it's not started yet", startCmd )
	subcmd.DefineParams( "appName" )

	subcmd = app.DefineSubCommand( "stop", "Stop <appName>, if it's running", stopCmd )
	subcmd.DefineParams( "appName" )

	subcmd = app.DefineSubCommand( "run", "Run <command> in <service> for <appName>", runCmd )
	subcmd.DefineParams( "appName", "service" )

	subcmd = app.DefineSubCommand( "rotate", "Perform a zero-downtime reload of <appName>", rotateCmd )
	subcmd.DefineParams( "appName" )

	subcmd = app.DefineSubCommand( "env", "Display env for <appname>", envCmd )
	subcmd.DefineParams( "appName" )

	subcmd = app.DefineSubCommand( "env:set", "Set an env variable for <appname>", envSetCmd )
	subcmd.DefineParams( "appName", "variable", "value" )

	subcmd = app.DefineSubCommand( "env:del", "Remove env variable for <appname>", envDelCmd )
	subcmd.DefineParams( "appName", "variable" )

	subcmd = app.DefineSubCommand( "letsencrypt", "Generate or renew ssl certificate for <appName>", letsencryptCmd )
	subcmd.DefineParams( "appName" )

	app.DefineSubCommand( "implode", "Remove leo from this server", implodeCmd )
}

/*
 * Let's do some sanity check before risking messing up with the system
 */
func checkConfig() ( err error ) {
	if len( os.Args ) > 1 && os.Args[1] == "setup" {
		if err = validator.TestConfigSetup() ; err != nil { return }
	} else if len( os.Args ) > 1 && os.Args[1] == "implode" {
		if err = validator.TestRoot() ; err != nil { return }
	} else {
		if err = validator.TestConfigRun() ; err != nil { return }
	}

	return
}

func main() {
	if err := checkConfig() ; err != nil {
		fmt.Println( err )
		os.Exit(1)
	}

	app.Start()
}
