package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/webornottoweb/ahri-log/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	}
}
