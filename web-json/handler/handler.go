package handler

import (
	"net/http"

	"github.com/seedboxtech/golang-training/web-json/reply"
)

type JsonHandler struct {
	name string // used to store state of handler
}

// internal method not exported
func (h *JsonHandler) setName(n string) {
	h.name = n
}

// external method. satisfies http.Handler interface
func (h *JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		h.setName(name) // set name

		reply.JSON(w, reply.StatusResponse{
			Status:  http.StatusOK,
			Message: "Message set successfully!",
		})
		return
	}

	reply.JSON(w, reply.NameResponse{
		Name: h.name,
	})
}
