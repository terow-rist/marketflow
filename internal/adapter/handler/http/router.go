package http

import "net/http"

func NewRouter(tempHandler *Temp) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(tempHandler.HandleTemp))

	return mux
}
