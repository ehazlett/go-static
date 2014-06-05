package main

import (
	"flag"
	"net/http"
)

var (
	STATIC_DIR string
)

func init() {
	flag.StringVar(&STATIC_DIR, "static-dir", "./", "Static directory to serve")
	flag.Parse()
}

func main() {
	panic(http.ListenAndServe(":3000", http.FileServer(http.Dir(STATIC_DIR))))
}
