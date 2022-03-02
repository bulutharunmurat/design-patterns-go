package adapter

import "fmt"

type WebClient struct {
	WebRequster WebRequester
}

type ClientReqObject struct {
	name string
	id   string
}

func (w *WebClient) DoWork() {
	clientObject := w.makeClientObject("req_name", "req_id")
	status := w.WebRequster.makeRequest(clientObject)

	if status == 200 {
		fmt.Println("request succeed")
	} else {
		fmt.Println("request NOT succeed")
	}
}

func (w *WebClient) makeClientObject(name, id string) ClientReqObject {
	clientReqObject := ClientReqObject{
		name: name,
		id:   id,
	}
	return clientReqObject
}
