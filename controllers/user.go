package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FelixMH/ProComments/commons"
	"github.com/FelixMH/ProComments/configuration"
	"github.com/FelixMH/ProComments/models"
)

// Login es el controlador de Login
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	db := configuration.GetConnection()
	if err != nil {
		log.Fatal("Error de conexión")
	}
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)

	db.Where("email = ? and password = ?", user.Email, pwd).First(&user)
	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)

		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatal("Error con el token")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "Usuario o Clave no válido",
			Code:    http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}
}

// UserCreate permite registrar un Usuario.
func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	if user.Password != user.ConfirmPassword {
		m.Message = "Las contraseñas no coinciden"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)
	user.Password = pwd

	pictureEmail := md5.Sum([]byte(user.Email))
	pic := fmt.Sprintf("%x", pictureEmail)
	user.Picture = "https://gravatar.com/avatar/" + pic + "?s=100"

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al registrarse: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Usuario Creado Con Éxito"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)

}
