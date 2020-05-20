package rpc

import (
	"errors"
	"rpc_blog_client/consul"
)

func SetUp() error {

	setUpTag("localhost:8081")
	adds:=consul.ConsulSetUp()
	if len(adds) <=0 {
		return errors.New("can't resolver article service")
	}
	setUpArticle(adds[0])
	return nil
}