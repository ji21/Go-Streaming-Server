package server

import (
	"io"
	"net/http"
)

//create a room and return a roomID
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

//join room and return sucess/failure
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {

}
