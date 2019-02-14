package discovery

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func NewConsulClient(hostname string) (*api.Client, error) {
	consulConfig := &api.Config{}
	consulConfig.Address = hostname

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}

	return consulClient, nil
}

func RegisterService(consulClient *api.Client, cfg *config.Config) error {
	err := consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      getServiceID(cfg.Consul.ServiceName, cfg.Server.Host),
		Name:    cfg.Consul.ServiceName,
		Address: cfg.Server.Host,
		Port:    cfg.Server.Port,
	})

	if err != nil {
		return err
	}

	return nil
}

func getServiceID(serviceName, hostname string) string {
	return fmt.Sprintf("%s-%s", serviceName, hostname)
}
