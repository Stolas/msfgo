package main

import (
	"fmt"
	"github.com/AtraxResearch/msfgo"
)

func main() {

	var conn msf.MetasploitConnection
	conn.Url = "http://192.168.178.36"
	conn.Port = 55553

	sess, err := msf.NewSession(conn, "msf", "msf")
	if err != nil {
		panic(err)
	}

	var all []msf.Console
	all, err = msf.ListConsoles(sess)
	if err != nil {
		panic(err)
	}

	fmt.Println("listConsole", all)

	msf.TerminateOwnSession(sess)
	return
}
