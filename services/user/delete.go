package user

import (
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Delete()")
}

func Deletes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Deletes()")
}
