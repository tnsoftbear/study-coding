package cache

import (
	"crypto/md5"
	"customer/internal/infra/storage"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const KEY_TPL = "cache:%s:%s"

type RedisCacher struct {
	rdb *redis.Client
	ns 	string
	exp time.Duration
}

func NewRedisCacher(namespace string, expiration time.Duration) *RedisCacher {
	return &RedisCacher{
		rdb: storage.NewRedisClient(),
		ns: namespace,
		exp: expiration,
	}
}

func (rc *RedisCacher) Set(key string, value interface{}) error {
	internalKey := rc.makeKey(key)
	set := rc.rdb.Set(internalKey, value, rc.exp)
	if set.Err() != nil {
		return set.Err()
	}
	return nil
}

func (rc *RedisCacher) Get(key string) (interface{}, error) {
	internalKey := rc.makeKey(key)
	get := rc.rdb.Get(internalKey)
	if get.Err() != nil {
		return "", get.Err()
	}
	return get.String(), nil
}

func (rc *RedisCacher) makeKey(key string) string {
	hashedKey := rc.hash(key)
	return fmt.Sprintf(KEY_TPL, rc.ns, hashedKey)
}

func (rc *RedisCacher) hash(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	output := hex.EncodeToString(hashBytes)
	return output
}
