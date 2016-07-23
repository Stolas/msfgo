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

func Save(session MetasploitSession) (err error) {
	m, err := SessionRequest(session, []string{"core.save"})

	err = getError(m)
	if err != nil {
		return
	}
}

func SetGlobal(session MetasploitSession, name string, value string) (err error) {
	m, err := SessionRequest(session, []string{"core.setg", name, value})

	err = getError(m)
	if err != nil {
		return
	}
}

func UnsetGlobal(session MetasploitSession, name string) (err error) {
	m, err := SessionRequest(session, []string{"core.unsetg", name})

	err = getError(m)
	if err != nil {
		return
	}
}

func Stop(session MetasploitSession) (err error) {
	m, err := SessionRequest(session, []string{"core.stop"})

	err = getError(m)
	if err != nil {
		return
	}
}
