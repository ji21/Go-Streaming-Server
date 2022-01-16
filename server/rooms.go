package server

import (
	"sync"

	"github.com/gorilla/websocket"
)

type User struct {
	host bool
	conn *websocket.Conn
}

type RoomMap struct {
	mutex sync.RWMutex
	rooms map[string][]User
}

//init room map struct
func (r *RoomMap) Init() {
	r.rooms = make(map[string][]User)
}
