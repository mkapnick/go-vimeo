package cache

import (
	"github.com/go-redis/redis"
)

// cache caches the frames using redis
// specs: 64MB cache size
var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	println(pong, err)
}

// Set a key/value pair in the redis store
func Set(key string, value string) (bool, error) {
	err := client.Set(key, value, 0).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}

// Get a value from the store associated with a key
func Get(key string) (string, error) {
	val, err := client.Get(key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

// Get info associated with redis db
func Info(key string) string {
	info := client.Info(key)
	return info.Val()
}
