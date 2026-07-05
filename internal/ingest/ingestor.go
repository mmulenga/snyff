package ingest

import (
	"io"
	"log"
	"net/http"
	"net/netip"
	"snyff/internal/store"
	"time"
)

type Ingestor struct {
	store store.RequestRepository
}

func NewIngestor(s store.RequestRepository) *Ingestor {
	return &Ingestor{s}
}

func (i *Ingestor) IngestHandler(w http.ResponseWriter, req *http.Request) {
	buffer := make([]byte, 8192) // 8KB buffer
	body := make([]byte, 0)
	bodySize := 0

	for {
		bytes, err := req.Body.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}

		body = append(body, buffer[:bytes]...)
		bodySize += bytes
	}
	defer req.Body.Close()

	addr, err := netip.ParseAddrPort(req.RemoteAddr)

	if err != nil {
		log.Println(err)
	}

	r := store.Request{
		Method: req.Method,
		Path: req.URL.Path,
		Query: req.RequestURI,
		Headers: req.Header,
		Body: body,
		Body_size_bytes: bodySize,
		Body_truncrated: false,
		Content_type: req.Header.Get("Content-Type"),
		Source_ip: addr.Addr(),
		Received_at: time.Now(),
	}

	if err := i.store.Save(&r); err != nil {
		log.Fatal(err)
	}
}
