package msf

import (
	"bytes"
	"gopkg.in/vmihailenco/msgpack.v2"
	"net/http"
	"strconv"
)

type Metasploit struct {
	Url           string
	Port          int32
	SessionTokens []string
}

func (msf *Metasploit) PackAndSend(input []string) map[string]string {
	client := &http.Client{
	// TODO: Allow ignore HTTPs
	}

	blob, err := msgpack.Marshal(input)
	if err != nil {
		panic(err)
	}

	resp, err := client.Post(msf.Url+":"+strconv.Itoa(int(msf.Port))+"/api/1.1/", "binary/message-pack", bytes.NewBuffer(blob))
	if err != nil {
		panic(err)
	}

	// TODO:
	// 200: The request was successfully processed
	// 500: The request resulted in an error
	// 401: The authentication credentials supplied were not valid
	// 403: The authentication credentials supplied were not granted access to the resource
	// 404: The request was sent to an invalid URL

	decodedMap := make(map[string]string)
	decoder := msgpack.NewDecoder(resp.Body)
	n, err := decoder.DecodeMapLen()
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			panic(err)
		}

		value, err := decoder.DecodeString()
		if err != nil {
			panic(err)
		}
		decodedMap[key] = value
	}

	return decodedMap
}
