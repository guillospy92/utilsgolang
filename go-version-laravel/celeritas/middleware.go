package celeritas

import (
	"net/http"
)

// SessionLoad add middleware sessions load
func (a *Accelerator) SessionLoad(next http.Handler) http.Handler {
	return a.Session.LoadAndSave(next)
}
