package name

import (
	"net/http"

	"github.com/gorilla/mux"
)

func View(w http.ResponseWriter, r *http.Request) {
	if name, ok := mux.Vars(r)["PARAM"]; ok {
		w.Write([]byte("Hello, " + name + "!"))
	}
	return
}
