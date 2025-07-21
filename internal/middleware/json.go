package middleware

import "net/http"

// JsonHeader устанавливает json хедер на ручку
func (m *Middleware) JsonHeader(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
