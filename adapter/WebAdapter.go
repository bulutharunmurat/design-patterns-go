package adapter

type WebAdapter struct {
	webService WebService
}

type ServiceReqObject struct {
	req_name string
	req_id   string
}

func (w *WebAdapter) Connect(webservice WebService) {
	w.webService = webservice
}

func (w *WebAdapter) makeRequest(c ClientReqObject) int {

	serviceReqObject := w.convert(c)

	status := w.webService.request(serviceReqObject)

	if len(status) > 0 {
		return 200
	} else {
		return 500
	}
}

func (w *WebAdapter) convert(c ClientReqObject) ServiceReqObject {
	serviceReqObject := ServiceReqObject{
		req_name: c.name,
		req_id:   c.id,
	}
	return serviceReqObject
}
