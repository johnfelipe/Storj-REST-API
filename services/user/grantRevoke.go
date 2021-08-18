package user

import (
	"fmt"
	"net/http"
)

func Grant(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Grant()")
}

func Revoke(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Revoke()")
}