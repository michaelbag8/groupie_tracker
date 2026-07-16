package middleware

import (
	"groupie_tracker/handlers"
	"net/http"
)

//Middleware to check for method instead of repeating it in all handlers
func RequireGet(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			handlers.RenderError(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r) //next(w,r) will still work

	}
}