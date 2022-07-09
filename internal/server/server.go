package server

import (
	"encoding/json"
	"net/http"
	"restapi/internal/database"
	"restapi/internal/models"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	router *mux.Router
	db     *database.Database
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
		db:     &database.Database{},
	}
}

func (s *Server) Start() error {
	s.configureRouter()
	if err := s.configureDatabase(); err != nil {
		return err
	}

	log.Info("Successfully opened database.")
	log.Info("Starting server...")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureRouter() {
	log.Info("Configure routes...")

	s.router.HandleFunc("/get/{key}", s.fetchValue()).Methods("GET")
	s.router.HandleFunc("/set", s.setRecord()).Methods("POST")
}

func (s *Server) configureDatabase() error {
	db, err := database.Open()
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	payloadBody, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Write(payloadBody)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func (s *Server) fetchValue() http.HandlerFunc {
	var (
		key, value string
		err        error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		key = mux.Vars(r)["key"]
		value, err = s.db.Get(key)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "not found")
			return
		}

		respondWithJson(w, http.StatusOK, value)
	}
}

func (s *Server) setRecord() http.HandlerFunc {
	var (
		decoder *json.Decoder
		data    *models.Record
		err     error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		decoder = json.NewDecoder(r.Body)
		if err = decoder.Decode(&data); err != nil {
			log.Error(err)
			respondWithError(w, http.StatusBadRequest, "invalid post scheme")
			return
		}

		log.Info("Set ", data)
		err = s.db.Set(data.Key, data.Value, time.Duration(data.Expiration))
		if err != nil {
			log.Error(err)
			respondWithJson(w, http.StatusBadRequest, "can't set values")
			return
		}

		respondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
	}
}
