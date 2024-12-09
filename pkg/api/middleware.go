package api

import (
	"net/http"

	"github.com/isnastish/nibble/pkg/log"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		log.Logger.Info("addr: %s path: %s method: %s", req.RemoteAddr, req.RequestURI, req.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(respWriter, req)
	})
}
