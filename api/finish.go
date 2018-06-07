package api

import (
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type Err struct {
	StatusCode int
	Message    string
}

func finish(w http.ResponseWriter, StatusCode int, Message string) {
	b, err := ffjson.MarshalFast(Err{StatusCode: StatusCode, Message: Message})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500"))
	}
	w.WriteHeader(StatusCode)
	w.Write(b)
}
