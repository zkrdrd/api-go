package refill

import (
	"api-go/constants"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Refill struct {
	Refill_ID   string `json:"id"`
	Customer_ID string `json:"CustomerFrom"`
	ATM         string `json:"ATM"`
	Amount      int    `json:"Amount"`
}

func (refill Refill) CreateRefill(responseWriter http.ResponseWriter, request *http.Request) {
	var UnixTimeCreate = int(time.Now().UnixNano())
	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(request.Body).Decode(&refill)
	if refill.Customer_ID == "" || refill.ATM == "" || refill.Amount <= 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(fmt.Sprintf(`HTTP status code %v returned!`, http.StatusBadRequest)))
	} else {
		refill.Refill_ID = strconv.Itoa(rand.Intn(UnixTimeCreate-constants.UnixTimeStart) + constants.UnixTimeStart)
		json.NewEncoder(responseWriter).Encode(refill)
	}
}
