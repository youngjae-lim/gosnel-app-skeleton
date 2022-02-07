package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/youngjae-lim/gosnel"
)

// Handlers is the type for handlers, and gives access to Gosnel and models
type Handlers struct {
	App    *gosnel.Gosnel
	Models data.Models
}

// Home is the handler to render the home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
