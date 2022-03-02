package adapter

type WebRequester interface {
	makeRequest(c ClientReqObject) int
}
