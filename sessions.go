package msf

import (
	"errors"
	"fmt"
)

type MetasploitSession struct {
	Connection   *MetasploitConnection
	SessionToken string
}

// NewSession creates a new session, the metasploit API uses the session for almost all functions.
func NewSession(conn MetasploitConnection, name string, password string) (session *MetasploitSession, err error) {
	m, err := conn.PackAndSend([]string{"auth.login", name, password})
	if err != nil {
		return
	}

	session = new(MetasploitSession)
	session.Connection = &conn

	result, err := getString(m, "result")
	if err != nil {
		return
	}

	if result == "success" {
		var token string
		token, err = getString(m, "token")

		session.SessionToken = token
	} else {
		err = errors.New(fmt.Sprintf("Invalid result: %s", result))
	}

	return
}

func TerminateSession(sess *MetasploitSession, LogoutToken string) {
	sess.Connection.PackAndSend([]string{"auth.logout", sess.SessionToken, LogoutToken})
}

func TerminateOwnSession(sess *MetasploitSession) {
	TerminateSession(sess, sess.SessionToken)
}

func SessionRequest(sess MetasploitSession, input []string) (m interface{}, err error) {
	input = append(input, sess.SessionToken)
	return sess.Connection.PackAndSend(input)
}
