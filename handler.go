package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func RedirectHandler(service LinkService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		url, err := service.Get(slug)
		if err == ErrDecodeFailure {
			msg := http.StatusText(http.StatusBadRequest)
			http.Error(w, msg, http.StatusBadRequest)
		} else if err == ErrNotFound {
			http.NotFound(w, r)
		} else {
			// `err` must be `nil`...
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		}
	})
}

func CreateLinkHandler(service LinkService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil || r.Body == http.NoBody {
			msg := http.StatusText(http.StatusBadRequest)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		payload := Link{}
		body, err := io.ReadAll(r.Body) // DOS attack vector
		defer r.Body.Close()
		if err == nil {
			err = json.Unmarshal(body, &payload)
		}
		if err != nil {
			msg := http.StatusText(http.StatusBadRequest)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		url := strings.TrimSpace(payload.URL)
		link, err := service.Create(url)
		if err != nil {
			msg := http.StatusText(http.StatusInternalServerError)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		res, _ := json.Marshal(link)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}
