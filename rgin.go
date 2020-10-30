package RGin

import (
	"log"
	"net/http"
)

type REngine struct {
	Router
}

func (r *REngine) ServeHTTP(w http.ResponseWriter, r2 *http.Request) {
	// url := r2.URL
	w.Write(([]byte)("http server"))
}

func (r *REngine) Run() {
	s := &http.Server{
		Addr:    ":8688",
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())
}
