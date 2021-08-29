package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type IResponse interface {
	GetData() interface{}
	GetStatusCode() int
	GetHeaders() map[string]string
}

func WriteResponse(res IResponse, w http.ResponseWriter) {
	for key, value := range res.GetHeaders() {
		w.Header().Add(key, value)
	}

	w.WriteHeader(res.GetStatusCode())
	bs, err := json.Marshal(res.GetData())
	if err != nil {
		log.Printf("Writing response: %v\n", err)
		return
	}

	_, _ = w.Write(bs)
}
