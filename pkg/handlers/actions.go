package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Session *gocql.Session
}

type Action struct {
	MessageID string    `json:"message_id,-,omitempty"`
	UserID    string    `json:"user_id,-,omitempty"`
	Status    string    `json:"status,-,omitempty"`
	Timestamp time.Time `json:"timestamp,-,omitempty"`
}

func (s *Server) SaveActionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer r.Body.Close()

	action := &Action{}
	err = json.Unmarshal(body, action)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	// TODO
	query := s.Session.Query("INSERT INTO tracking (messageID, userID, status, timestamp) VALUES (?, ?, ?, ?);",
		action.MessageID, action.UserID, action.Status, time.Now())
	err = query.Exec()
	if err != nil {
		w.WriteHeader(500)
		log.Printf("failed SQL insert: %v\n", err)
		return
	}
	w.WriteHeader(201)
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Write([]byte(`{"result":"ok"}`))
}

func (s *Server) GetActionStatusHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	messageID := queryString.Get("message_id")
	userID := queryString.Get("user_id")

	fmt.Printf("message_id: %v, user_id: %v\n", messageID, userID)
	// TODO
	var status string
	query := s.Session.Query("SELECT status FROM tracking WHERE messageID = ? AND userID = ?;", messageID, userID)
	err := query.Scan(&status)
	if err == gocql.ErrNotFound {
		w.WriteHeader(404)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		log.Printf("failed SQL query: %v\n", err)
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Write([]byte(fmt.Sprintf(`{"result":"%s"}`, status)))
}
