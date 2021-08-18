package user

import (
	"fmt"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Upload()")
}

func Download(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Download()")
}

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from List()")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Delete()")
}
