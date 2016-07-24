package msf

import (
	"errors"
)

func getMap(m interface{}) (dict map[interface{}]interface{}, err error) {
	dict, correct := m.(map[interface{}]interface{})
	if !correct {
		err = errors.New("Invalid Type")
		return
	}
	return
}

func getString(m interface{}, key string) (str string, err error) {
	dict, err := getMap(m)
	if err != nil {
		return
	}

	buff := dict[key].([]uint8)
	str = string(buff)
	return
}

func getInt(m interface{}, key string) (i int, err error) {
	// var str string
	// str, err = getString(m, key)

	// if err != nil {
	// 	return
	// }
	//fmt1.Println(str)
	//i, err = strconv.Atoi(str)
	dict, err := getMap(m)
	if err != nil {
		return
	}

	i64 := dict[key].(int64)
	i = int(i64) // TODO: Check for data loss.
	return
}

func getBool(m interface{}, key string) (v bool, err error) {
	dict, err := getMap(m)
	if err != nil {
		return
	}

	if dict[key] != nil {
		v = dict[key].(bool)
	}
	return
}

func getError(m interface{}) (err error) {
	// This needs to be rewriten, it is a total mess.
	if m == nil {
		err = errors.New("Empty interface")
		return
	}

	dict, err := getMap(m)
	if err != nil {
		return
	}

	if _, ok := dict["error"]; ok {
		isError, _ := getBool(m, "error")
		if !isError {
			return
		}
	}

	if _, ok := dict["result"]; ok {
		var result string
		result, _ = getString(m, "result")
		if result == "failure" {
			err = errors.New("Failed to write to Console")
			return
		}
	}

	if _, ok := dict["error_string"]; ok {
		errorStr, _ := getString(m, "error_string")
		err = errors.New(errorStr)
	}

	return
}
