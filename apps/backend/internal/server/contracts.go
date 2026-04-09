package server

import "net/http"

type Server interface {
	Start() error
	createHttpServer() *http.Server
}
