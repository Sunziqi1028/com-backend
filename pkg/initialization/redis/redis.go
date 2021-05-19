package redis

import "github.com/gotomicro/ego-component/eredis"

var Client *eredis.Component

func Init() (err error) {
	Client = eredis.Load("redis.stub").Build()
	return
}
