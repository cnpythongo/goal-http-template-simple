package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var captchaRedisExpiration = 600 * time.Second // 15 minutes

type CaptchaRedisStore struct {
	client     *redis.Client // redis连接
	expiration time.Duration // 有效时长
	keyPrefix  string        // 存储键的前缀
}

func NewCaptchaRedisStore(client *redis.Client, keyPrefix string) *CaptchaRedisStore {
	return &CaptchaRedisStore{
		client:     client,
		expiration: captchaRedisExpiration,
		keyPrefix:  keyPrefix,
	}
}

func (c *CaptchaRedisStore) SetExpiration(expiration time.Duration) {
	c.expiration = expiration
}

func (c *CaptchaRedisStore) Set(id string, value string) error {
	err := c.client.Set(context.Background(), c.keyPrefix+id, value, c.expiration).Err()
	return err
}

func (c *CaptchaRedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	var value string
	var err error
	if clear {
		value, err = c.client.GetDel(ctx, c.keyPrefix+id).Result()
	} else {
		value, err = c.client.Get(ctx, c.keyPrefix+id).Result()
	}
	if err != nil {
		return ""
	}
	return value
}

func (c *CaptchaRedisStore) Verify(id, answer string, clear bool) (match bool) {
	return c.Get(id, clear) == answer
}
