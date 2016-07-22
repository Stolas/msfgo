package msf

import "testing"

func TestFullFlow(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	t.Log("Starting Full Test")

	var conn MetasploitConnection
	conn.Url = "192.168.178.36"
	conn.Port = 55554

	t.Log("Creating Session")
	session_1, _ := NewSession(conn, "msf", "msf")
	ver1, _, _ := GetVersion(*session_1)

	t.Log("Creating Session")
	session_2, _ := NewSession(conn, "msf", "msf")
	ver2, _, _ := GetVersion(*session_2)

	if ver1 != ver2 {
		t.Errorf("ver1 and ver2 do not match")
	}
}
