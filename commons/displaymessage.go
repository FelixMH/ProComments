package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FelixMH/ProComments/models"
)

// DisplayMessage devolverá un mensaje al cliente
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(m.Code)
	w.Write(j)

}
