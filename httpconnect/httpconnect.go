// Package http-connect
// 
// Contains all the functions to connect to the server, contains the all thing into
// an object

package httpconnect

import "net/http"
import "fmt"
import "bytes"

func ConnectTo(server string, secure bool) string {
	var (
		request		string
		respHTTP	*http.Response
		err		error
	)
	// Connect to the server 
	if secure == true {
		request="https://"+server+"/"
	} else {
		request="http://"+server+"/"
	}
	respHTTP, err = http.Get(request)
	if err != nil {
		fmt.Printf("Ouch!\n")
	}
	resp := new(bytes.Buffer)
	resp.ReadFrom(respHTTP.Body)
	return resp.String()
}
