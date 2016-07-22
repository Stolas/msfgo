package msf

func (msf *Metasploit) GetVersion() (version string, ruby string) {
	return msf.GetVersionById(0)
}

func (msf *Metasploit) GetVersionById(index int) (version string, ruby string) {
	session := msf.SessionTokens[index]
	//if session == nil {
	//	return nil, nil
	//}
	return msf.GetVersionBySession(session)
}

func (msf *Metasploit) GetVersionBySession(SessionToken string) (version string, ruby string) {
	versionmap := msf.PackAndSend([]string{"core.version", SessionToken})
	return versionmap["version"], versionmap["ruby"]
}
