package main

import (
	_ "github.com/impactmarketingspecialists/registrator/consul"
	_ "github.com/impactmarketingspecialists/registrator/consulkv"
	_ "github.com/impactmarketingspecialists/registrator/etcd"
	_ "github.com/impactmarketingspecialists/registrator/skydns2"
	_ "github.com/impactmarketingspecialists/registrator/zookeeper"
)
