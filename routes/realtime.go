package routes

import (
	"github.com/gorilla/mux"
	"github.com/olahol/melody"

	"net/http"
)

// SetRealtimeRouter: Ruta para el realtime
func SetRealtimeRouter(router *mux.Router) {
	m := melody.New()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
}
