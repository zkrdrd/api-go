package transfer

import "context"

type CacheIn struct {
	Account string `json:"account"`
	Amount  string `json:"amount"`
}

type Transfer struct {
	Account_from string `json:"account_from"`
	Account_to   string `json:"account_to"`
	Amount       string `json:"amount"`
}

type Transaction struct {
}

func NewBalance() *Transaction {
	return &Transaction{}
}

func (b *Transaction) CacheIn(ctx context.Context, indata []byte, money *CacheIn) ([]byte, error) {
	return indata, nil
}

func (b *Transaction) Transfer(ctx context.Context, indata []byte, money *Transfer) ([]byte, error) {
	return indata, nil
}

// func (addition Add_money) CreateAddition(responseWriter http.ResponseWriter, request *http.Request) {
// 	var UnixTimeCreate = int(time.Now().UnixNano())
// 	responseWriter.Header().Set("Content-Type", "application/json")
// 	_ = json.NewDecoder(request.Body).Decode(&addition)
// 	if addition.addition_ID == "" || addition.ATM == "" || addition.Amount <= 0 {
// 		responseWriter.WriteHeader(http.StatusBadRequest)
// 		responseWriter.Write([]byte(fmt.Sprintf(`HTTP status code %v returned!`, http.StatusBadRequest)))
// 	} else {
// 		addition.addition_ID = strconv.Itoa(rand.Intn(UnixTimeCreate-constants.UnixTimeStart) + constants.UnixTimeStart)
// 		json.NewEncoder(responseWriter).Encode(addition)
// 	}
// }

// func (remittance Remittance) CreateRemittance(responseWriter http.ResponseWriter, request *http.Request) {
// 	var UnixTimeCreate = int(time.Now().UnixNano())
// 	responseWriter.Header().Set("Content-Type", "application/json")
// 	_ = json.NewDecoder(request.Body).Decode(&remittance)
// 	if remittance.Customer_From == "" ||
// 		remittance.Customer_To == "" ||
// 		remittance.Amount <= 0 {
// 		responseWriter.WriteHeader(http.StatusBadRequest)
// 		responseWriter.Write([]byte(fmt.Sprintf(`HTTP status code %v returned!`, http.StatusBadRequest)))
// 	} else {
// 		remittance.Remittance_ID = strconv.Itoa(rand.Intn(UnixTimeCreate-constants.UnixTimeStart) + constants.UnixTimeStart)
// 		json.NewEncoder(responseWriter).Encode(remittance)
// 	}

// }
