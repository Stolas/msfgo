package msf

import (
	"errors"
	"strconv"
)

type Console struct {
	Sess   *MetasploitSession
	Id     int
	Prompt string
	Busy   bool
}

func NewConsole(session *MetasploitSession) (console *Console, err error) {
	m, err := SessionRequest(*session, []string{"console.create"})

	err = getError(m)
	if err != nil {
		return
	}

	console = new(Console)
	console.Sess = session

	var id string
	id, err = getString(m, "id")
	if err != nil {
		return
	}

	console.Id, err = strconv.Atoi(id)
	if err != nil {
		return
	}

	console.Prompt, err = getString(m, "prompt")
	if err != nil {
		return
	}

	console.Busy, err = getBool(m, "busy")
	if err != nil {
		return
	}

	return
}

func ListConsoles(session *MetasploitSession) (console []Console, err error) {
	_, err = SessionRequest(*session, []string{"console.list"})
	panic("TODO -- Implement")
	// for k, v := range m {
	// 	fmt.Printf("key=%v, value=%v", k, v)
	// }
	return
}

func (console *Console) Read(p []byte) (n int, err error) {
	m, err := SessionRequest(*console.Sess, []string{"console.read", strconv.Itoa(console.Id)})
	// TODO: Now it gets called 10000 times?
	// TODO: Check EOF

	console.Prompt, err = getString(m, "prompt")
	if err != nil {
		return
	}

	console.Busy, err = getBool(m, "busy")
	if err != nil {
		return
	}

	var str string
	str, err = getString(m, "data")
	p = []byte(str)

	if len(p) == 0 {
		err = errors.New("EOF")
	}

	return
}

func (console *Console) Write(p []byte) (n int, err error) {
	m, err := SessionRequest(*console.Sess, []string{"console.write", strconv.Itoa(console.Id), string(p)})

	err = getError(m)
	if err != nil {
		return
	}

	n, err = getInt(m, "wrote")
	return
}
