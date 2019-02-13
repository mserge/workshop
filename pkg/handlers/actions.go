package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Server struct{}

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

	w.WriteHeader(201)
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Write([]byte(`{"result":"ok"}`))
}

func (s *Server) GetActionStatusHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	messageID := queryString.Get("message_id")
	userID := queryString.Get("user_id")

	fmt.Printf("message_id: %v, user_id: %v\n", messageID, userID)

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Write([]byte(fmt.Sprintf(`{"result":"%s"}`, "")))
}
