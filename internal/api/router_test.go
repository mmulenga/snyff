package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFindRequestHandler(t *testing.T) {
	http.NewRequest("GET", "/requests", nil)
	res, err := http.Get("http://localhost:8080/requests/019fb0fb-39be-7437-9323-e649ea65fcc5")
	if err != nil {
		t.Skip()
	}
	defer res.Body.Close()

	bytes := make([]byte, 8192)
	n, err := res.Body.Read(bytes)
	if err != nil {
		t.Skip()
	}

	fmt.Println(n)
	fmt.Println(string(bytes))
}
