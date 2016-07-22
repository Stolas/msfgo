package msf

import "errors"

type MetasploitSession struct {
	Connection   *MetasploitConnection
	SessionToken string
}

// NewSession creates a new session, the metasploit API uses the session for almost all functions.
func NewSession(conn MetasploitConnection, name string, password string) (session *MetasploitSession, err error) {
	loginresult, err := conn.PackAndSend([]string{"auth.login", name, password})
	if err != nil {
		return
	}

	session = new(MetasploitSession)
	session.Connection = &conn

	if loginresult["result"] == "success" {
		session.SessionToken = loginresult["result"]
	} else {
		err = errors.New("Failed to login")
	}

	return
}

func TerminateSession(sess *MetasploitSession, LogoutToken string) {
	sess.Connection.PackAndSend([]string{"auth.logout", sess.SessionToken, LogoutToken})
}

func TerminateOwnSession(sess *MetasploitSession) {
	TerminateSession(sess, sess.SessionToken)
}

func SessionRequest(sess MetasploitSession, input []string) (decodedMap map[string]string, err error) {
	input = append(input, sess.SessionToken)
	return sess.Connection.PackAndSend(input)
}
