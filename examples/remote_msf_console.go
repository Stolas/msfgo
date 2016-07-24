package main

import (
	"bufio"
	"bytes"
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

	var console *msf.Console
	console, err = msf.NewConsole(sess)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(console)

	buf := new(bytes.Buffer)
	fmt.Println("######################################")
	buf.ReadFrom(console)
	fmt.Println(buf.String())

	n, err := writer.WriteString("help")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", n)

	msf.TerminateOwnSession(sess)
}
