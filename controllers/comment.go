package controllers

import (
	"net/http"
	"github.com/FelixMH/ProComments/models"
	"encoding/json"
	"fmt"
	"github.com/FelixMH/ProComments/commons"
	"github.com/FelixMH/ProComments/configuration"
)

// CreateComment: crea un comentario.

func CreateComment(w http.ResponseWriter, r *http.Request)  {
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al crear el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado con éxito."
	commons.DisplayMessage(w, m)


}