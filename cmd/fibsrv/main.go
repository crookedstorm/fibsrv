package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/crookedstorm/fibsrv/pkg/api"
)

// func init() {

// }
func main() {
	r, fibs := api.InitRouter()
	answer1, _ := fibs.LookupFib(9)
	answer2, _ := fibs.BigFib(173)
	fmt.Printf("You got a low of %d\n", answer1)
	fmt.Printf("You got a high of %s\n", answer2.String())
	// r.Run(":3000")
	log.Printf("[info] start http server listening on :3000")
	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
