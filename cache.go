package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

type BaseCache struct {
	Key string `json:"key"`
}

type SetCache struct {
	BaseCache
	Value  string `json:"value"`
	Expire int64  `json:"expire"`
}

type IntCache struct {
	BaseCache
	Step int64 `json:"step"`
}

var c *cache.Cache

func init() {
	c = cache.New(cache.DefaultExpiration, cache.NoExpiration)
}

func CacheAction(action, argumentStr string) (interface{}, error) {

	switch strings.ToUpper(action) {
	case "DEL":
		return cacheDel(argumentStr)
	case "GET":
		return cacheGet(argumentStr)
	case "SET":
		return cacheSet(argumentStr)
	case "INCR":
		return cacheIncr(argumentStr)
	case "DECR":
		return cacheDecr(argumentStr)
	}
	return "", errors.New("Unsupported command ")

}

func cacheGet(arg string) (string, error) {
	var bc BaseCache
	err := json.Unmarshal([]byte(arg), &bc)
	if err != nil {
		return "", err
	}
	foo, found := c.Get(bc.Key)
	if found {
		return fmt.Sprintf("%s", foo), nil
	}
	return "", errors.New("Not Found ")
}

func cacheSet(arg string) (string, error) {
	sc := SetCache{Expire: -1}
	err := json.Unmarshal([]byte(arg), &sc)
	if err != nil {
		return "", err
	}
	if sc.Key == "" {
		return "", errors.New("Key is empty ")
	}
	c.Set(sc.Key, sc.Value, time.Duration(sc.Expire)*time.Second)
	return sc.Key, nil
}

func cacheDel(arg string) (string, error) {
	var bc BaseCache
	err := json.Unmarshal([]byte(arg), &bc)
	if err != nil {
		return "", err
	}
	if bc.Key == "" {
		return "", errors.New("Key is empty ")
	}
	c.Delete(bc.Key)
	return "", nil
}

func cacheIncr(arg string) (string, error) {

}

func cacheDecr(arg string) (string, error) {

}
