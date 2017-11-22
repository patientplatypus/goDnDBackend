package users

import (
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Instance(r)

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
