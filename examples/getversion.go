package main

import (
	"fmt"
	"github.com/AtraxResearch/msfgo"
)

func main() {
	var conn msf.MetasploitConnection
	conn.Url = "http://192.168.178.36"
	conn.Port = 55553

	session, err := msf.NewSession(conn, "msf", "msf")
	if err != nil {
		panic(err)
	}

	msfver, rubyver, err := msf.GetVersion(*session)
	if err != nil {
		panic(err)
	}

	fmt.Println("Metasploit Version:", msfver)
	fmt.Println("Ruby Version:", rubyver)
}
