package user

import (
	"fmt"
	"net/http"
)


func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Update()")
}




