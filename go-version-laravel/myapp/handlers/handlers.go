package handlers

import (
	"net/http"
	
	"github.com/CloudyKit/jet/v6"
	"github.com/guillospy92/utilsgolang/go-version-laravel/celeritas"
)

type Handlers struct {
	App *celeritas.Accelerator
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering")
	}
}

func (h *Handlers) HomeWithGo(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) SessionPage(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "foo", "bar")

	value := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", value)

	err := h.App.Render.JetPage(w, r, "session", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering")
	}
}
