package main

import (
	"gomagru/auth"
	"net/http"
)

func main() {
	auth.GetHandler()
	http.ListenAndServe(":4000", nil)
}
