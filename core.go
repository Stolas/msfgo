package msf

func GetVersion(session MetasploitSession) (version string, ruby string, err error) {
	m, err := SessionRequest(session, []string{"core.version"})

	err = getError(m)
	if err != nil {
		return
	}

	version, err = getString(m, "version")
	if err != nil {
		return
	}

	ruby, err = getString(m, "ruby")
	return
}
