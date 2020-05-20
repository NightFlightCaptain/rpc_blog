package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type RegisterInfo struct {
	Host           string
	Port           int
	ServiceName    string
	UpdateInterval time.Duration
}


type Register struct {
	Target string //consul的地址
	Ttl    int
}

func NewConsulRegister(target string, ttl int) *Register {
	return &Register{Target: target, Ttl: ttl}
}

func generateServiceId(name, host string, port int) string {
	return fmt.Sprintf("%s-%s-%d", name, host, port)
}

func (c *Register) Register(info RegisterInfo) error {
	config := consulapi.DefaultConfig()
	config.Address = c.Target
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("create consul client error:", err.Error())
		return err
	}

	serviceId := generateServiceId(info.ServiceName, info.Host, info.Port)

	reg := &consulapi.AgentServiceRegistration{
		ID:      serviceId,
		Name:    info.ServiceName,
		Tags:    []string{info.ServiceName},
		Port:    info.Port,
		Address: info.Host,
	}

	if err = client.Agent().ServiceRegister(reg); err != nil {
		panic(err)
	}

	check := consulapi.AgentServiceCheck{
		TTL:    fmt.Sprintf("%ds", c.Ttl),
		Status: consulapi.HealthPassing,
	}
	err = client.Agent().CheckRegister(
		&consulapi.AgentCheckRegistration{
			ID:                serviceId,
			Name:              info.ServiceName,
			ServiceID:         serviceId,
			AgentServiceCheck: check,
		},
	)

	if err != nil {
		return fmt.Errorf("initial register service check to consul error: %s", err.Error())
	}
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		x := <-ch
		log.Println("receive signal: ", x)
		// un-register service
		c.DeRegister(info)

		s, _ := strconv.Atoi(fmt.Sprintf("%d", x))
		os.Exit(s)
	}()

	go func() {
		ticker := time.NewTicker(info.UpdateInterval)
		for {
			<-ticker.C
			err = client.Agent().UpdateTTL(serviceId, "", check.Status)
			if err != nil {
				log.Println("update ttl of service error: ", err.Error())
			}
		}
	}()

	return nil
}

func (c *Register) DeRegister(info RegisterInfo) error {

	serviceId := generateServiceId(info.ServiceName, info.Host, info.Port)

	config := consulapi.DefaultConfig()
	config.Address = c.Target
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("create consul client error:", err.Error())
	}

	err = client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		log.Println("deregister service error: ", err.Error())
	} else {
		log.Println("deregistered service from consul server.")
	}

	err = client.Agent().CheckDeregister(serviceId)
	if err != nil {
		log.Println("deregister check error: ", err.Error())
	}

	return nil
}
