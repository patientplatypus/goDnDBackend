package character

import (
	"fmt"
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func ClassType(w http.ResponseWriter, r *http.Request) {
	log.Print("inside Class")
	class := r.FormValue("class")
	fmt.Fprintln(w, "Your Class: ", class)
	log.Print("Value of r in Class: ", r)
	log.Print("Value of r==nil: ", r == nil)
	// log.Print("Value of store: ", store)
	session, _ := config.Instance(r)
	log.Print("inside Class 2")
	log.Print("value of session: ", session)
	log.Print("value of session.Values['authenticated']", session.Values["authenticated"])
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Print("Verboten!")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}
