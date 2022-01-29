package server

import (
	"io"
	"net/http"
)

//create a room and return a roomID
func FetchNewsRequestHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
