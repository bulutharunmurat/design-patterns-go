package main

import (
	"design-patterns-go/adapter"
)

func main() {

	host := "https://google.com"
	service := adapter.WebService{Host: host}
	a := adapter.WebAdapter{}
	a.Connect(service)

	client := adapter.WebClient{WebRequster: &a}

	client.DoWork()

}
