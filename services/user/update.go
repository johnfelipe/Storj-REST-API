package user

import (
	"fmt"
	"net/http"
)

// Fetch the object by id
// Delete it
// Add new updates at the same id
// 1. Download
// 2. Delete
// 3. Upload

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Update()")
}




