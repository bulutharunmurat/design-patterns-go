package adapter

type WebService struct {
	Host string
}

func (w *WebService) request(s ServiceReqObject) string {
	return "string response from web service, and host is " + w.Host
}
