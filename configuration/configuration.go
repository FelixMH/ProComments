package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// Función que obtiene los datos de conexión a base de datos, retorna una estructura como respuesta y ésta tiene
// los datos que precisamente necesita nuestro servidor para conectarse a BD.
func GetConfiguration() Configuration {
	var c Configuration

	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c

}

// Función que obtiene la conexión a la base de datos con GORM, ORM con distintas Bases de datos y lenguajes del mismo.
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
