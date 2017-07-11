package config

func AppDir( appName string ) string {
	return LeoDir() + "/apps/" + appName
}
