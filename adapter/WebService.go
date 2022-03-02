package adapter

import "fmt"

type WebService struct {
	Host string
}

func (w *WebService) request(s ServiceReqObject) string {
	fmt.Println("request maded")
	return "string response from web service, and host is " + w.Host
}
