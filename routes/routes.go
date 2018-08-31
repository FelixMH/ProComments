package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes Inicia todas las rutas de la api
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)
	SetVoteRouter(router)

	return router
}
