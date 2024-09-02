package remittance

import (
	"api-go/constants"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Remittance struct {
	Remittance_ID string `json:"id"`
	Customer_From string `json:"CustomerFrom"`
	Customer_To   string `json:"CustomerTo"`
	Amount        int    `json:"Amount"`
}

func (remittance Remittance) CreateRemittance(responseWriter http.ResponseWriter, request *http.Request) {
	var UnixTimeCreate = int(time.Now().UnixNano())
	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(request.Body).Decode(&remittance)
	if remittance.Customer_From == "" ||
		remittance.Customer_To == "" ||
		remittance.Amount <= 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(fmt.Sprintf(`HTTP status code %v returned!`, http.StatusBadRequest)))
	} else {
		remittance.Remittance_ID = strconv.Itoa(rand.Intn(UnixTimeCreate-constants.UnixTimeStart) + constants.UnixTimeStart)
		json.NewEncoder(responseWriter).Encode(remittance)
	}

}
