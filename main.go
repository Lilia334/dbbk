package main

import (
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pg"
		"log"
	"net/http"
) 

func main() {
	initDB()
	initController()

log.Fatal(http.ListenAndServe(":9000", router))
}
