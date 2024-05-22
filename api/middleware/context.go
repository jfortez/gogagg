package middleware

import (
	"context"
	"net/http"
)

type contextHandler struct {
	ctx context.Context
	h   http.Handler
}

func (c contextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.h.ServeHTTP(w, r.WithContext(c.ctx))
}

// New
func NewContextHandler(ctx context.Context) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return contextHandler{ctx, next}
	}
}
