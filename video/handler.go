package video

import (
	"fmt"
	"net/http"
)

func VideoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/video" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "This is a VideoHandler.")
}
