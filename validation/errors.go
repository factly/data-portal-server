package validation

import (
	"net/http"
	"strings"

	"github.com/factly/data-portal-server/util"
)

// InvalidID - response for invalid ID
func InvalidID(w http.ResponseWriter, r *http.Request) {
	var msg []string
	msg = append(msg, "Invalid id")
	util.Render(w, http.StatusBadRequest, msg)
}

// RecordNotFound - response for record not found
func RecordNotFound(w http.ResponseWriter, r *http.Request) {
	var msg []string
	msg = append(msg, "Record not found")
	util.Render(w, http.StatusBadRequest, msg)
}

// ValidErrors - errors from validator
func ValidErrors(w http.ResponseWriter, r *http.Request, msg string) {
	err := strings.Split(msg, "\n")
	util.Render(w, http.StatusBadRequest, err)
}
