package character

import (
	"fmt"
	"net/http"
)

func BackgroundType(w http.ResponseWriter, r *http.Request) {
	background := r.FormValue("background")
	fmt.Fprintln(w, "Your Background: ", background)
}
