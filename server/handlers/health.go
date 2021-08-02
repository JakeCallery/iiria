package handlers

import (
	"log"
	"net/http"
)

type Health struct {
	l *log.Logger
}

func NewHealth(l *log.Logger) *Health {
	return &Health{l}
}

func (h *Health) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("We are up!\n"))
}
