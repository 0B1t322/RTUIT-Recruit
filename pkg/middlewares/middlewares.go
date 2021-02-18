package middlewares

import "net/http"

// TODO checkAuth

// ContentTypeJSONMiddleware set header content-type to application json
func ContentTypeJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	})
}