package routes

import (
	"github.com/FelixMH/ProComments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

/* SetCommentRouter: define la ruta de comentarios para la api */
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("", controllers.CreateComment).Methods("POST")
	subRouter.HandleFunc("/", controllers.CommentGetAll).Methods("GET")

	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(controllers.ValidateToken),
		negroni.Wrap(subRouter),
	),
	)
}
