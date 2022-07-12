package internal

import (
	"github.com/nsmithuk/local-kms/src/data"
	"net/http"
)

// HandleRequest wraps handleRequest() so it can be exported.
func HandleRequest(w http.ResponseWriter, r *http.Request, database *data.Database) {
	handleRequest(w, r, database)
}
