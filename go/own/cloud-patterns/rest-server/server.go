/**
 * curl -X "PUT" -d "KeyB data" -v localhost:8080/v1/key-b
 * curl localhost:8080/v1/key-b
 * curl localhost:8080/v1/list
 * curl -X "DELETE" localhost:8080/v1/key-b
 */
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrorNoSuchKey = errors.New("no such key")

func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Put(key string, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()
	return nil
}

func Delete(key string) error {
	_, err := Get(key)
	if err != nil {
		return err
	}
	store.Lock()
	delete(store.m, key)
	store.Unlock()
	return nil
}

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func keyValueListHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("%v", store))
	w.WriteHeader(http.StatusOK)
}

func keyValueGetHanlder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(value))
	w.WriteHeader(http.StatusOK)
}

func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	err := Delete(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, fmt.Sprintf("Successfully deleted key: %v", key))
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/list", keyValueListHandler)
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueGetHanlder).Methods("GET")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
