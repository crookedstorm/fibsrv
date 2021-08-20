package main

import (
	"log"
	"net/http"
	"time"

	"github.com/crookedstorm/fibsrv/pkg/api"
)

// func init() {

// }
func main() {
	r := api.InitRouter()

	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("[info] start http server listening on :3000")
	log.Fatal(s.ListenAndServe())
}
