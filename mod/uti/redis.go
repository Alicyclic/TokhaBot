package uti

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
)

func NewDatabase(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}

func CreateHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	key = fmt.Sprintf("%x", hash.Sum(nil))
	return key
}

func (db *Database) SetIfNotExists(key string, value string) (bool, error) {
	if err := db.Client.SetNX(CreateHash(key), value, time.Minute+3).Err(); err != nil {
		defer db.Client.Close()
		return false, err
	}
	return true, nil
}

func (db *Database) Get(key string) (string, error) {
	return db.Client.Get(CreateHash(key)).Result()
}
