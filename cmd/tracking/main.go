package main

import (
	"fmt"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/config"
	"gitlab.k8s.gromnsk.ru/workshop/montalcini/pkg/tracking"
)

func main() {
	fmt.Println("Tracking workshop")
	cfg, err := config.GetConfig()
	fmt.Printf("%+v", cfg)
	if err != nil {
		fmt.Println(err)
	}
	tracking.Run(cfg)
}
