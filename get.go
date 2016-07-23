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
	isError, _ := getBool(m, "error")
	if !isError {
		return
	}

	errorStr, _ := getString(m, "error_string")
	err = errors.New(errorStr)

	return
}
