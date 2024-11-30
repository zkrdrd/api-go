//go:build integration

package tests

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/zkrdrd/api-go/pkg/server"
)

const (
	addressServer = `127.0.0.1:8080`
)

func TestAccounting(t *testing.T) {
	//mux := http.NewServeMux()

	ctx := context.Background()
	//usersService := users.NewUsers()
	//transactionAddService := transfer.NewBalance()

	//handler := usersService.Handlers(mux)

	server := server.NewServer(addressServer)
	//server.AddHandler(usersService.Handlers())
	//server.AddHandler(transactionAddService.Handlers())
	server.Run(ctx)

	time.Sleep(time.Second * 1)

	// Prepare message
	dataForCheck := map[string]string{
		//`/users`:               `{"id": "1", "UserName": "asdf", "Password": "asdfasdf"}`,
		`/transaction/CacheIn`:  `{"id": "1", "from": "123", "to": "1, "amount": 50}`,
		`/transaction/Transfer`: `{"id": "1", "from": "1", "to": "2", "amount": 50}`,
	}

	buf := &bytes.Buffer{}

	for key, value := range dataForCheck {
		buf.WriteString(value)

		// request builder
		req, err := http.NewRequest(http.MethodPost, `http://`+addressServer+key, buf)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		// do request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		// check status
		if res.StatusCode != http.StatusOK {
			t.Error(`invalid code: `, res.StatusCode)
			t.Fail()
			log.Print(key)
		}

		// check response data
		resData, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if string(resData) != value {
			t.Error(`invalid response `, string(resData))
			t.Fail()
		}
	}

}
