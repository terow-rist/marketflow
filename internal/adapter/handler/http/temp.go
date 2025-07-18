package http

import (
	"fmt"
	"net/http"
)

type Temp struct{}

func NewTemp() *Temp {
	return &Temp{}
}

func (t *Temp) HandleTemp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jopa")
}
