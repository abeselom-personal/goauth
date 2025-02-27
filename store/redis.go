package store

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
}

func (db *RedisDB) Connect(addr string) error {
	db.client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := db.client.Ping(context.Background()).Result()
	return err
}

func (db *RedisDB) Disconnect() error {
	return db.client.Close()
}
