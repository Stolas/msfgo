package msf

import (
	"github.com/ugorji/go/codec"
	"io"
	"net/http"
	"strconv"
)

type MetasploitConnection struct {
	Url  string
	Port int32
}

func (conn MetasploitConnection) PackAndSend(input []string) (decodedMap map[string]string, err error) {
	client := &http.Client{
	// TODO: Allow ignore HTTPs
	}

	r, w := io.Pipe()
	var h codec.Handle = new(codec.MsgpackHandle)
	var enc *codec.Encoder = codec.NewEncoder(w, h)

	err = enc.Encode(input)
	if err != nil {
		return
	}

	port := strconv.Itoa(int(conn.Port))
	_, err = client.Post(conn.Url+":"+port+"/api/1.1/", "binary/message-pack", r)
	if err != nil {
		return
	}

	// TODO:
	// 200: The request was successfully processed
	// 500: The request resulted in an error
	// 401: The authentication credentials supplied were not valid
	// 403: The authentication credentials supplied were not granted access to the resource
	// 404: The request was sent to an invalid URL

	return
}
