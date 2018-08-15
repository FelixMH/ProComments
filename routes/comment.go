package routes

import (
	"github.com/gorilla/mux"
	"github.com/FelixMH/ProComments/controllers"
	"github.com/urfave/negroni"
)

/* SetCommentRouter: define la ruta de comentarios para la api */
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CreateComment).Methods("POST")

	router.PathPrefix(prefix).Handler(negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}