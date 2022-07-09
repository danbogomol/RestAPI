package database

import (
	"time"

	"github.com/go-redis/redis"
)

type Database struct {
	client *redis.Client
}

func Open() (*Database, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "database:6379",
		Password: "",
		DB:       0,
	})

	err := db.Ping().Err()
	if err != nil {
		return nil, err
	}

	return &Database{client: db}, nil
}

func (db *Database) Get(key string) (string, error) {
	value, err := db.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (db *Database) Set(key, value string, expiration time.Duration) error {
	err := db.client.Set(key, value, expiration*time.Millisecond).Err()
	if err != nil {
		return err
	}

	return nil
}