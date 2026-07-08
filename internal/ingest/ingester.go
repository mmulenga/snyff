package ingest

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/netip"
	"snyff/internal/store"
	"time"
)

type Ingester struct {
	requestStore store.RequestRepository
}

func NewIngester(s store.RequestRepository) *Ingester {
	return &Ingester{s}
}

func (i *Ingester) IngestHandler(w http.ResponseWriter, req *http.Request) {
	request := store.Request{}
	reader := http.MaxBytesReader(w, req.Body, 1024 * 1024)
	buffer := make([]byte, 8192) // 8KB buffer
	body := make([]byte, 0)
	bodySize := 0

	for {
		bytes, err := reader.Read(buffer)
		body = append(body, buffer[:bytes]...)
		bodySize += bytes

		if err != nil {
			if err == io.EOF {
				break
			} 
			var e *http.MaxBytesError
        	if errors.As(err, &e) {
				log.Println(err)
				request.Body_truncated = true
				break
        	}
			break
		}
	}
	defer req.Body.Close()

	addr, err := netip.ParseAddrPort(req.RemoteAddr)

	if err != nil {
		log.Println(err)
	}

	request.Method = req.Method
	request.Path = req.URL.Path
	request.Query = req.RequestURI
	request.Headers = req.Header
	request.Body = body
	request.Body_size_bytes = bodySize
	request.Content_type = req.Header.Get("Content-Type")
	request.Source_ip = addr.Addr()
	request.Received_at = time.Now()

	if err := i.requestStore.Save(&request); err != nil {
		log.Fatal(err)
	}
}
