package tracking

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/config"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/discovery"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/handlers"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/storage"
	"log"
	"net/http"
)

type Server struct {
}

func Run(cfg *config.Config) {

	client, err := discovery.NewConsulClient(cfg.Consul.Hostport)
	if err != nil {
		log.Fatalf("failed to init  consul client: %v\n", err)
	}
	serviceNode, err := discovery.GetServiceNode(client, "cassandra")
	if err != nil {
		log.Fatalf("failed to get live db: %v\n", err)
	}
	log.Printf(" Going to connect %s:%d", serviceNode.Address, serviceNode.Port)

	keyspace := cfg.Storage.Keyspace
	//hostport := fmt.Sprintf("%s:%d", cfg.Storage.Host, cfg.Storage.Port)
	hostport := fmt.Sprintf("%s:%d", serviceNode.Address, serviceNode.Port)
	session, err := storage.InitStorage(hostport, keyspace)
	if err != nil {
		log.Fatalf("failed to init  storage: %v\n", err)
	}

	server := handlers.Server{Session: session}

	err = discovery.RegisterService(client, cfg)
	if err != nil {
		log.Fatalf("failed to register service: %v\n", err)
	}

	defer func() {
		err := discovery.DeregisterService(client, cfg)
		log.Printf("failed to deregister service: %v\n", err)
	}()

	handlers.RegisterMetrics()

	http.HandleFunc("/save", server.SaveActionHandler)
	http.HandleFunc("/get-status", server.GetActionStatusHandler)

	http.HandleFunc("/healthz", server.HealthHandler)
	http.HandleFunc("/readyz", server.ReadyzHandler)
	http.HandleFunc("/", server.ReadyzHandler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Service starting on port %d..\n", cfg.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
	if err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
