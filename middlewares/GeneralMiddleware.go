package middlewares

import "net/http"

func GeneralMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/html;charset=UTF-8")
		h.ServeHTTP(w, r)
	})
}