package main

import (
	"fmt"
	"github.com/AtraxResearch/msfgo"
	"time"
)

func SpawnSession(conn msf.MetasploitConnection) (sess *msf.MetasploitSession) {
	sess, err := msf.NewSession(conn, "msf", "msf")
	if err != nil {
		panic(err)
	}
	fmt.Println("Session Started:", sess.SessionToken)
	return
}

func KillSession(sess msf.MetasploitSession) {
	fmt.Println("Killing Session:", sess.SessionToken)
	msf.TerminateOwnSession(&sess)
}

func SessionLifeCycle(conn msf.MetasploitConnection) {
	sess := SpawnSession(conn)
	time.Sleep(500)
	KillSession(*sess)
}

func main() {
	var conn msf.MetasploitConnection
	conn.Url = "http://192.168.178.36"
	conn.Port = 55553

	SessionLifeCycle(conn)
	fmt.Println("done")
}
