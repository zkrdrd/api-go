package users

import "context"

type Users struct {
	User_ID   string `json:"id"`
	User_Name string `json:"UserName"`
	Password  string `json:"Password"`
}

func NewUsers() *Users {
	return &Users{}
}

func (b *Users) Create(ctx context.Context, indata []byte) ([]byte, error) {
	return indata, nil
}

// func (cust handlers.Customer) CreateCustomer(responseWriter http.ResponseWriter, request *http.Request) {
// 	var UnixTimeCreate = int(time.Now().UnixNano())
// 	responseWriter.Header().Set("Content-Type", "application/json")
// 	_ = json.NewDecoder(request.Body).Decode(&cust)
// 	if cust.User_Name == "" || cust.Password == "" {
// 		responseWriter.WriteHeader(http.StatusBadRequest)
// 		responseWriter.Write([]byte(fmt.Sprintf(`HTTP status code %v returned!`, http.StatusBadRequest)))
// 	} else {
// 		cust.Customer_ID = strconv.Itoa(rand.Intn(UnixTimeCreate-constants.UnixTimeStart) + constants.UnixTimeStart)
// 		json.NewEncoder(responseWriter).Encode(cust)
// 	}
// }
