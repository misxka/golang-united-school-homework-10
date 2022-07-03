package headers

import (
	"net/http"
	"strconv"
	"strings"
)

func View(w http.ResponseWriter, r *http.Request) {
	valueA, errA := strconv.Atoi(r.Header.Get("A"))
	if errA != nil {
		panic(errA)
	}
	valueB, errB := strconv.Atoi(r.Header.Get("B"))
	if errB != nil {
		panic(errB)
	}

	headerName := "a+b"

	w.Header().Set(strings.ToLower(headerName), strconv.Itoa(valueA+valueB))
	return
}
