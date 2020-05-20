package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net"
	"strconv"
)

func ConsulSetUp() []string {
	var lastIndex uint64
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8500"

	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return nil
	}

	services, metainfo, err := client.Health().Service("article", "article", true, &consulapi.QueryOptions{
		WaitIndex: lastIndex,
	})

	if err != nil {
		log.Printf("error retrieving instances from Consul: %v", err)
	}
	lastIndex = metainfo.LastIndex

	adds := make([]string,0)
	for _, service := range services {
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		adds = append(adds, net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port)))
	}
	return adds
}
