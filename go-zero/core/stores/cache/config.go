package cache

import "github.com/wenj91/mctl/go-zero/core/stores/redis"

type (
	ClusterConf []NodeConf

	NodeConf struct {
		redis.RedisConf
		Weight int `json:",default=100"`
	}
)
