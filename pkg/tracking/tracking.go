package tracking

import (
	"fmt"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/config"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/handlers"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {

	server := handlers.Server{}

	http.HandleFunc("/save", server.SaveActionHandler)
	http.HandleFunc("/get-status", server.GetActionStatusHandler)

	http.HandleFunc("/healthz", server.HealthHandler)

	fmt.Printf("Service starting on port %d..\n", cfg.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
	if err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
