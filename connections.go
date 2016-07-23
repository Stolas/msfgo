package msf

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ugorji/go/codec"
	"net/http"
	"strconv"
)

type MetasploitConnection struct {
	Url  string
	Port int32
}

func (conn MetasploitConnection) PackAndSend(input []string) (m interface{}, err error) {
	client := &http.Client{
	// TODO: Allow ignore HTTPs
	}

	b := new(bytes.Buffer)
	var h codec.Handle = new(codec.MsgpackHandle)
	var enc *codec.Encoder = codec.NewEncoder(b, h)

	err = enc.Encode(input)
	if err != nil {
		return
	}

	var resp *http.Response
	port := strconv.Itoa(int(conn.Port))
	resp, err = client.Post(fmt.Sprintf("%s:%s/api/1.1", conn.Url, port), "binary/message-pack", b)
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case 200:
		// 200: The request was successfully processed
		break
	case 500:
		// 500: The request resulted in an error
		err = errors.New("Server Error")
		return
	case 401:
		// 401: The authentication credentials supplied were not valid
		err = errors.New("Invalid Credentials")
		return
	case 403:
		// 403: The authentication credentials supplied were not granted access to the resource
		err = errors.New("The authentication credentials supplied were not granted access to the resource")
		return
	case 404:
		// 404: The request was sent to an invalid URL
		err = errors.New("Invalid URL")
		return
	default:
		break
	}

	var dec *codec.Decoder = codec.NewDecoder(resp.Body, h)
	err = dec.Decode(&m)
	return
}
