package app

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type contextKey int

const (
	traceID contextKey = iota
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func withMiddleware(handler http.HandlerFunc, middlewares ...middleware) http.HandlerFunc {
	if len(middlewares) == 0 {
		return handler
	}

	wrapped := handler

	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	return wrapped
}

func primeRequestContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), traceID, uuid.New())
		next(w, r.WithContext(ctx))
	}
}

func recoverWrap(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			e := recover()
			if e != nil {
				var err error
				switch t := e.(type) {
				case string:
					err = errors.New("💀 " + t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}

				uuid := r.Context().Value(traceID).(uuid.UUID)
				log.Println("Error for request: " + uuid.String())
				log.Println(err.Error())

				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		next(w, r)
	}
}
