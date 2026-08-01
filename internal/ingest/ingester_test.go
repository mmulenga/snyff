package ingest

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

type Charge struct {
	Id   *string `json:"id"`
	Type *string `json:"type"`
	Data Payment `json:"data"`
}

type Payment struct {
	Amount   int     `json:"amount"`
	Currency *string `json:"currency"`
}

func TestIngest(t *testing.T) {
	body := strings.NewReader(`{"id":"evt_1PabcXYZ","type":"charge.succeeded","data":{"amount":2000,"currency":"usd"}}`)
	mockRequest := httptest.NewRequest("GET", "/requests", body)
	mockRequest.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c := Charge{}
	request := ingest(w, mockRequest)
	if err := json.Unmarshal(request.Body, &c); err != nil {
		t.Fail()
	}

	if strings.Compare(*c.Data.Currency, "usd") != 0 {
		t.Fail()
	}
}
