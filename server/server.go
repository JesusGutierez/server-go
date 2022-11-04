package server

import "net/http"

type Country struct {
	Name     string
	Language string
}

func New(addr string) *http.Server {
	return &http.Server{
		Addr: addr,
	}
}
