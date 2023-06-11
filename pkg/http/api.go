package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Api struct {
	http.Handler
	Router *httprouter.Router
	Port   int
}

func (a *Api) Initialize() {
	a.Router = httprouter.New()
	a.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.Router.ServeHTTP(w, r)
	})

}

func (a *Api) Start() {
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", a.Port),
		Handler: a,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
