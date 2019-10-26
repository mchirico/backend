package api

import (
	"context"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
	File string
}

func NewAPI() *API {
	return &API{File: "/dummy.csv"}
}

func NewAPIfile(file string) *API {
	return &API{File: file}
}

func (api *API) MainListen(ctx context.Context) {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome v4"))
	})
	r.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(ReadFile(api.File)))
	})

	server := &http.Server{Addr: ":3000", Handler: r}
	//http.ListenAndServe(":3000", r)
	server.ListenAndServe()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error")
	}
}

func ReadFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
