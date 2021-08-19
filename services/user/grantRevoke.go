package user

import (
	"fmt"
	"net/http"
)

func Share(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Share()")
}

func Shares(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Shares()")
}

func RevokeGrant(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from RevokeGrant()")
}


func RevokeGrants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from RevokeGrants()")
}