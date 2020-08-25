package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye World")

	wr.Write([]byte("Byeeee"))
}
