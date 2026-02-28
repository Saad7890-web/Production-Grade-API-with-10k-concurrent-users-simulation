package main

import (
	"encoding/json"
	"net/http"
)

type API struct {
	store *Store
}

func NewAPI(store *Store) *API {
	return &API{store: store}
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (a *API) setHandler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	a.store.Set(req.Key, req.Value)
	w.WriteHeader(http.StatusCreated)
}

func (a *API) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	val, ok := a.store.Get(key)
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"value": val,
	})
}