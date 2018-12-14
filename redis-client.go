package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

type RedisCommon struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RedisGet struct {
	RedisCommon
	Key string `json:"key"`
}

type RedisSet struct {
	RedisCommon
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration int    `json:"expiration"`
}

var clients = make(map[string]redis.Client)

func (p RedisCommon) getClient() redis.Client {
	addr := fmt.Sprintf("%s:%d", p.Host, p.Port)
	client, m := clients[addr]
	if !m {
		fmt.Println(fmt.Sprintf("%s:%d", p.Host, p.Port))
		clients[addr] = *redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", p.Host, p.Port),
			Password: "",
			DB:       0,
			PoolSize: 4,
		})
		return clients[addr]
	}
	return client

}

func RedisAction(action, argumentStr string) (interface{}, error) {

	switch strings.ToUpper(action) {
	case "DEL":
		return redisDel(argumentStr)
	case "GET":
		return redisGet(argumentStr)
	case "SET":
		return redisSet(argumentStr)
	case "INCR":
		return redisIncr(argumentStr)
	case "HGETALL":
		return redisHgetall(argumentStr)
	}
	return "", errors.New("Unsupported command ")
}

func redisDel(argument string) (int64, error) {
	var p RedisGet
	err := json.Unmarshal([]byte(argument), &p)

	if err != nil {
		return 0, err
	}
	redisClient := p.getClient()

	keys := strings.Split(p.Key, " ")

	return redisClient.Del(keys...).Result()
}

func redisGet(argument string) (string, error) {
	var p RedisGet
	err := json.Unmarshal([]byte(argument), &p)

	if err != nil {
		return "", err
	}
	redisClient := p.getClient()
	return redisClient.Get(p.Key).Result()
}

func redisSet(argument string) (string, error) {
	var p RedisSet
	err := json.Unmarshal([]byte(argument), &p)

	if err != nil {
		return "", err
	}
	redisClient := p.getClient()
	return redisClient.Set(p.Key, p.Value, time.Duration(p.Expiration)).Result()
}

func redisIncr(argument string) (int64, error) {
	var p RedisGet
	err := json.Unmarshal([]byte(argument), &p)

	if err != nil {
		return 0, err
	}
	redisClient := p.getClient()
	return redisClient.Incr(p.Key).Result()
}

func redisHgetall(argument string) (map[string]string, error) {
	var p RedisGet
	err := json.Unmarshal([]byte(argument), &p)

	if err != nil {
		return nil, err
	}
	redisClient := p.getClient()
	return redisClient.HGetAll(p.Key).Result()
}
