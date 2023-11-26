package main

import (
    "net/http"
    "github.com/gorilla/mux"
)
type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

type APIServer struct{
    listenAddr string

}
    func NewApiServer(listenAddr string) *APIServer {
        return &APIServer{
            listenAddr:listenAddr,
        }
    }
    
    func (s *APIServer) Run() {
        router := mux.NewRouter()

        router.HandleFunc("/account", s.handleAccount).Methods("GET")
    }


    func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
        return nil

    }


    func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
        return nil

    }



    func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
        return nil

    }

    func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
        return nil

    }
    func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
        return nil

    }
