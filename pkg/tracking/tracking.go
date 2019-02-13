package tracking

import (
	"fmt"
	"log"
	"net/http"
)

func Run(config *config.Config) {

	http.HandleFunc("/save", server.SaveActionHandler)
	http.HandleFunc("/get-status", server.GetActionStatusHandler)

	http.HandleFunc("/healthz", server.HealthHandler)

	fmt.Printf("Service starting on port %d..\n", cfg.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
	if err != nil {
		log.Fatalf("failed")
	}
}
