package data

import (
	"io/ioutil"
	"net/http"
)

func View(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		w.Write(append([]byte("I got message:\n"), body...))
	}
	return
}
