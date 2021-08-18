package iam

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Login()")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Home()")
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Refresh()")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Register()")
}
