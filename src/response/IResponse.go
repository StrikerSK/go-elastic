package response

import (
	"encoding/json"
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
	bs, _ := json.Marshal(res.GetData())
	_, _ = w.Write(bs)
}
