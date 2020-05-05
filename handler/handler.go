package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dcaponi/echo/interfaces"
)

type Handler struct {
	Controller interfaces.Controller
}

func NewHandler(cont interfaces.Controller) Handler {
	return Handler{Controller: cont}
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.post(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) get(w http.ResponseWriter, r *http.Request) {
	// logic around unpacking a request and formatting it for a controller lives
	// here. Should only deal in request/response and byte[] really...
	id := r.URL.Query().Get("id")
	var (
		err  error
		i    int
		resp []byte
	)
	if id == "" {
		resp = h.Controller.Get()
	} else {
		if i, err = strconv.Atoi(id); err != nil {
			_, _ = w.Write([]byte("Invalid Query Params Given"))
			return
		}
		resp = h.Controller.GetOne(i)
	}
	_, _ = w.Write(resp)
}

func (h Handler) post(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	resp := h.Controller.Write(body)
	_, _ = w.Write(resp)
}
