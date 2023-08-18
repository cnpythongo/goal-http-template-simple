package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"

	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/status"
)

type processFunc func(string, []byte)

func ListenRedisList(queueName string, fn processFunc) {
	status.AddWaitGroup()
	defer status.DoneWaitGroup()
	for status.IsRunnning() {
		calList(queueName, fn)
	}
	return
}

func calList(queueName string, fn processFunc) {
	c := redisPool.Get()
	defer c.Close()
	reply, err := redis.ByteSlices(c.Do("BRPOP", queueName, 10))
	if err != nil {
		if err != redis.ErrNil {
			log.GetLogger().Error("BRPOP:", err)
		}
		return
	}
	if len(reply) == 2 {
		fn(queueName, reply[1])
	}
}

func PubRedisList(queueName string, payload interface{}) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.GetLogger().Error("PubRedisList Error:", err)
		return
	}
	_, err = LPush(queueName, bytes)
	if err != nil {
		log.GetLogger().Error("PubRedisList0 Error:", err)
	}
}
