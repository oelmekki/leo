package setup

import (
	"os"
	"fmt"
)

/*
 * If you change this, remember to change the sudo rule for it in README
 */
var ChownScriptPath = "/usr/local/bin/chown_leo_deploy"

/*
 * This script is needed to make sure leo can remove apps, without
 * needing to use dangerous sudo commands.
 *
 * Initially, this rule was added:
 *
 *     leo-deploy ALL=(ALL) NOPASSWD: /bin/rm -r /home/leo-deploy/apps/*
 *
 * But it would have allowed to do this:
 *
 *     ln -s /etc /home/leo-deploy/apps/my_app
 *     rm -r /home/leo-deploy/apps/my_app/
 *
 * ... and delete all system conf files.
 *
 * The script created here will be created as root, so user can't override
 * its content. It chowns the home directory, which user can't remove either.
 * All in all, it seems the safer way to allow user to reclaim ownership
 * of files in their home directory.
 *
 * Note that this has a severe implication, though : leo should not be used
 * with host directory mounted as volumes - user has to use data container.
 * If they use host directories, files in it will be chowned at an unpredicable
 * point, and app using the volume won't be able to delete those files anymore
 * if it's not ran as root. Use data containers.
 */
func SetupChownScript() ( err error ) {
	filename := ChownScriptPath
	fmt.Printf( "Installing %s...\n", filename )

	file, err := os.Create( filename )
	if err != nil { return err }

	_, err = file.WriteString( chownLeoDeployContent )
	if err != nil { return err }

	if err = file.Chmod( 0755 ) ; err != nil { return }

	file.Close()

	return
}

/*
 * Used when removing leo from system
 */
func removeChownScript() ( err error ) {
	if err = os.Remove( ChownScriptPath ) ; err != nil { return }
	return
}


var chownLeoDeployContent = `#!/usr/bin/env bash
chown -R leo-deploy:leo-deploy /home/leo-deploy
`
