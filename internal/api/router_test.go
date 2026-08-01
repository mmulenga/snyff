package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFindRequestHandler(t *testing.T) {
	res, err := http.Get("http://localhost:8080/requests/019fb0fb-39be-7437-9323-e649ea65fcc5")
	if err != nil {
		t.Skip()
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			t.Fail()
		}
	}()

	bytes := make([]byte, 8192)
	n, err := res.Body.Read(bytes)
	if err != nil {
		t.Skip()
	}

	fmt.Println(n)
	fmt.Println(string(bytes))
}
