package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/FelixMH/ProComments/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	// Se usa para validar el token.
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}

	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("No se puedo leer el archivo público")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privateKey", err)
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a PublicKey", err)
	}
}

// GenerateJWT Genera el token para el cliente.
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour * 364).Unix(),
			Issuer: "ProComments",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token.")
	}

	return result

}
