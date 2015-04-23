package glory

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goware/throttler"
)

const UpdateEndpoint = "/glory/update/available"

func Glorify(secret string, callback func()) {
	// Limit to only 1 update request at a time
	limit := throttler.Limit(1)
	http.Handle(UpdateEndpoint, limit(&updateHandler{secret: secret, callback: callback}))
}

type updateHandler struct {
	secret   string
	callback func()
}

func (uh *updateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		url := r.FormValue("url")
		sha1 := r.FormValue("sha1")
		timestamp := r.FormValue("timestamp")
		signature := r.FormValue("signature")

		fmt.Printf("URL: %s, sha1: %s, timestamp: %s, signature: %s\n", url, sha1, timestamp, signature)

		ur := NewUpdateRequest(url, sha1)
		ur.SetTimestamp(timestamp)

		if ur.Expired(time.Now().UTC()) {
			fmt.Println("Expired")
			http.Error(w, "Request Has Expired", http.StatusUnauthorized)
			return
		}

		if ur.Signature(uh.secret) != signature {
			fmt.Println("Bad Signature")
			http.Error(w, "Bad Signature", http.StatusUnauthorized)
			return
		}

		err := updateExecutable(url, sha1)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "Update Succeeded")
			uh.callback()
		}

	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
