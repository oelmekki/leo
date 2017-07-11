package main

import (
	"github.com/jwaldrip/odin/cli"
)

func defaultUsage( c cli.Command ) {
	c.Usage()
}
var app = cli.New( "1.0.0", "leo local client", defaultUsage )

/*
 * Usual way to setup odin's accepted commands
 */
func init() {
	subcmd := app.DefineSubCommand( "init", "Setup leo for app in current directory", initCmd )
	subcmd.DefineParams( "appName", "server" )

	app.DefineSubCommand( "deploy", "Deploy app from current directory to server", deployCmd )

	subcmd = app.DefineSubCommand( "run", "Run <command> for <service>", runCmd )
	subcmd.DefineParams( "service" )

	app.DefineSubCommand( "env", "Display current environment configuration", envCmd )

	subcmd = app.DefineSubCommand( "env:set", "set given variable to given value", envSetCmd )
	subcmd.DefineParams( "variable", "value" )

	subcmd = app.DefineSubCommand( "env:del", "remove given variable", envDelCmd )
	subcmd.DefineParams( "variable" )

	app.DefineSubCommand( "letsencrypt", "Generate or renew ssl certificate", letsencryptCmd )
}

func main() {
	app.Start()
}
