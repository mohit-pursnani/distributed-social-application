package raftHelper

/*
	Utilizing examples from https://pkg.go.dev/net/http to manipulate the raft links found in raftexamples and appending
	to those links to communicate with the raft. Raft example starts 3 raft nodes.

	"http://127.0.0.1:12380/" + key will mean kv pair will be stored in raft 1 and so on

	Below code uses the curl command to encode and decode key value pairs:
	eg: curl -L http://127.0.0.1:12380/my-key -XPUT -d hello => Encode
	and curl -L http://127.0.0.1:12380/my-key => Decode

	reference: https://mholt.github.io/curl-to-go/ & https://www.codershood.info/2017/06/25/http-curl-request-golang/
*/

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var mutex sync.Mutex

func encodeData(value interface{}) (string, error) {
	var encodeBuffer bytes.Buffer
	newEncoder := gob.NewEncoder(&encodeBuffer)
	errorVal := newEncoder.Encode(value)
	if errorVal != nil {
		return "", errorVal
	} else {
		//encodedData = string(buffer)
		return encodeBuffer.String(), nil
	}
}

func checkPUT(value interface{}, operation string) (string, error) {
	var encodedData string
	var errorVal error

	if operation == "PUT" {
		encodedData, errorVal = encodeData(value)
		if errorVal != nil {
			return "", errorVal
		}
	}
	return encodedData, errorVal
}

func decodedDataRes(response *http.Response) ([]byte, error) {
	var decodedData []byte

	respBody := response.Body

	decodedData, err := ioutil.ReadAll(respBody)

	if err != nil {
		log.Println("Error", err)
		return decodedData, err
	}

	return decodedData, nil
}

func RequestForService(key string, value interface{}, operation string) (string, error) {

	var encodedData string
	encodedData = ""

	var errorVal error

	//3 raft urls that will be started after goreman start
	raftServerURL1 := "http://127.0.0.1:12380/" + key
	raftServerURL2 := "http://127.0.0.1:22380/" + key
	raftServerURL3 := "http://127.0.0.1:32380/" + key

	var raftURLList [3]string
	raftURLList[0] = raftServerURL1
	raftURLList[1] = raftServerURL2
	raftURLList[2] = raftServerURL3

	encodedData, errorVal = checkPUT(value, operation)
	if errorVal != nil {
		log.Print("Error: ", errorVal)
		return "", errorVal
	}

	//var resFromRaft *http.Response
	resFromRaft := make(chan *http.Response)

	for i := 0; i < 3; i += 1 {
		var reader *strings.Reader
		var request *http.Request

		if value == nil {

		} else {
			reader = strings.NewReader(encodedData)
		}

		go func(index int) {
			mutex.Lock()

			if operation == "PUT" {
				request, errorVal = http.NewRequest("PUT", raftURLList[index], reader)
			} else if operation == "GET" {
				request, errorVal = http.NewRequest("GET", raftURLList[index], reader)
			}

			if errorVal != nil {
				log.Print("Error: ", errorVal)
				mutex.Unlock()
				return
			}

			var response *http.Response

			defClient := http.DefaultClient

			response, errorVal = defClient.Do(request)

			if errorVal != nil {
				mutex.Unlock()
				return
			}

			mutex.Unlock()

			resFromRaft <- response
		}(i)

	}

	response := <-resFromRaft

	decodedData, err := decodedDataRes(response)

	if err != nil {
		return "", err
	}

	dataToSend := string(decodedData)

	return dataToSend, nil
}
