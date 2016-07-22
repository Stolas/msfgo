package msf

func GetVersion(session MetasploitSession) (version string, ruby string, err error) {
	versionmap, err := SessionRequest(session, []string{"core.version"})

	version = versionmap["version"]
	ruby = versionmap["ruby"]
	return
}
