package redis

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal/pkg/common/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPool       *redis.Pool
	redisNotInitErr = errors.New("redis is not initialized.")
	KeyNotExist     = errors.New("key does not exist")
)

func Init(cfg *config.RedisConfig) error {
	url := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	redisPool = &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: 20 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				return nil, err
			}
			if cfg.Auth != "" {
				_, err = c.Do("AUTH", cfg.Auth)
				if err != nil {
					return nil, err
				}
			}
			_, err = c.Do("SELECT", cfg.Db)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return nil
}

func Get(key string) ([]byte, error) {
	if redisPool == nil {
		return nil, redisNotInitErr
	}
	c := redisPool.Get()
	defer c.Close()

	reply, err := c.Do("GET", key)
	if err != nil {
		return nil, err
	}
	val, ok := reply.([]byte)
	if ok {
		return []byte(val), nil
	}
	return val, KeyNotExist
}

func Set(key string, ttl int64, value []byte) error {
	if redisPool == nil {
		return redisNotInitErr
	}
	c := redisPool.Get()
	defer c.Close()
	if _, err := c.Do("set", key, value); err != nil {
		return err
	}
	if ttl > 0 {
		if _, err := c.Do("expire", key, ttl); err != nil {
			return err
		}
	}
	return nil
}

func Del(key string) error {
	if redisPool == nil {
		return redisNotInitErr
	}
	c := redisPool.Get()
	defer c.Close()
	_, err := c.Do("del", key)
	return err
}

func Exist(key string) bool {
	if redisPool == nil {
		panic(nil)
	}
	c := redisPool.Get()
	defer c.Close()

	exist, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		panic(err)
	}
	return exist
}

func HSet(key string, values map[string][]byte, ttl int64) error {
	if redisPool == nil {
		panic(nil)
	}
	c := redisPool.Get()
	defer c.Close()
	var redisValues []interface{}
	redisValues = append(redisValues, key)
	for k, v := range values {
		redisValues = append(redisValues, k)
		redisValues = append(redisValues, v)
	}
	_, err := c.Do("HMSET", redisValues...)
	if err == nil && ttl > 0 {
		if _, err := c.Do("expire", key, ttl); err != nil {
			return err
		}
	}
	return err
}

func HGet(key string, vkeys ...string) (map[string][]byte, error) {
	if redisPool == nil {
		panic(nil)
	}
	c := redisPool.Get()
	defer c.Close()

	result := make(map[string][]byte)
	var redisKeys []interface{}
	redisKeys = append(redisKeys, key)
	for _, vk := range vkeys {
		redisKeys = append(redisKeys, vk)
	}
	values, err := redis.Values(c.Do("HMGET", redisKeys...))
	if err != nil {
		return result, err
	}
	for i := 0; i < len(vkeys); i++ {
		if values[i] == nil {
			return result, KeyNotExist
		}
		result[vkeys[i]] = values[i].([]byte)
	}
	return result, nil
}

func LPush(key string, args ...[]byte) (int64, error) {
	if redisPool == nil {
		panic(nil)
	}
	c := redisPool.Get()
	defer c.Close()
	vs := []interface{}{}
	vs = append(vs, key)
	for _, v := range args {
		vs = append(vs, v)
	}
	reply, err := redis.Int64(c.Do("LPUSH", vs...))
	return reply, err
}

func Increment(key string) error {
	if redisPool == nil {
		panic(nil)
	}
	c := redisPool.Get()
	defer c.Close()
	_, err := c.Do("INCR", key)
	return err
}

func BRPop(queueName string) ([][]byte, error) {
	if redisPool == nil {
		return nil, redisNotInitErr
	}
	c := redisPool.Get()
	defer c.Close()

	reply, err := redis.ByteSlices(c.Do("BRPOP", queueName, 5))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Close() error {
	if redisPool != nil {
		err := redisPool.Close()
		redisPool = nil
		return err
	}
	fmt.Println("redis close done.")
	return nil
}
