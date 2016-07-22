package msf

import "fmt"

// Login sends the auth.login msgpack to metasploit.
func (msf *Metasploit) Login(name string, passwd string) {
	loginmap := msf.PackAndSend([]string{"auth.login", name, passwd})
	if loginmap["result"] == "success" {
		msf.SessionTokens = append(msf.SessionTokens, loginmap["token"])
	} else {
		fmt.Println("w00p Failure!")
	}
}

func (msf *Metasploit) Logout() {
	sessionToken := msf.SessionTokens[0]
	msf.LogoutToken(sessionToken, sessionToken)
}

func (msf *Metasploit) LogoutToken(SessionToken string, LogoutToken string) {
	msf.PackAndSend([]string{"auth.logout", SessionToken, LogoutToken})
	// TODO: Remove the old element.
}
