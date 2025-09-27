package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mugen64/turtlor/pkg/apperrors"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteErrorResponse(w http.ResponseWriter, err error) error {
	if errors.Is(err, apperrors.ApiError{}) {
		be, ok := err.(apperrors.ApiError)
		if !ok {
			bep, ok := err.(*apperrors.ApiError)
			if !ok {
				fmt.Println("Failed to convert error to ApiError", err)
				return WriteJSON(w, http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
			}
			be = *bep
		}

		return WriteJSON(w, be.Status, be)
	}

	fmt.Println("Internal Server Error ", err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error"})
}

type ApiHandlerFunc func(http.ResponseWriter, *http.Request) error

func (h ApiHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		WriteErrorResponse(w, err)
	}
}

func IsDebugHeaderSet(r *http.Request) bool {
	return r.Header.Get("Debug") == "reveal-body-logs"
}
