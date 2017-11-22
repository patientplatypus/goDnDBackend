package character

import (
	"fmt"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func RaceType(w http.ResponseWriter, r *http.Request) {
	race := r.FormValue("race")
	sqlConfig := config.Sql_config()
	fmt.Fprintln(w, "Your Race: ", race)
	fmt.Fprintln(w, "Sql_config: ", sqlConfig)
	fmt.Fprintln(w, "Yolo")
}
