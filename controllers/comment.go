package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/FelixMH/ProComments/commons"
	"github.com/FelixMH/ProComments/configuration"
	"github.com/FelixMH/ProComments/models"
)

// CreateComment: crea un comentario.

func CreateComment(w http.ResponseWriter, r *http.Request) {
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

// CommentGetAll: Obtiene todos los comentarios.
func CommentGetAll(w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	// vote := models.Vote{}

	r.Context().Value(&user)
	// Todo lo que llega despues del ? en la api.
	vars := r.URL.Query()

	db := configuration.GetConnection()
	defer db.Close()

	cComment := db.Where("parent_id = 0")
	if order, ok := vars["order"]; ok {
		if order[0] == "votes" {
			cComment = cComment.Order("votes desc, created_at desc")
		}
	} else {
		if idlimit, ok := vars["idlimit"]; ok {
			registerByPage := 30
			offset, err := strconv.Atoi(idlimit[0])
			if err != nil {
				log.Println("Error:", err)
			}
			cComment = cComment.Where("id BETWEEN ? AND ?", offset-registerByPage, offset)
		}
		cComment = cComment.Order("id desc")
	}

	cComment.Find(&comments)

	j, err := json.Marshal(comments)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir los comentarios en json"
		commons.DisplayMessage(w, m)
		return
	}

	if len(comments) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m.Code = http.StatusNoContent
		m.Message = "No hay contenido de comentarios"
		commons.DisplayMessage(w, m)
	}

}
