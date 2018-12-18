package main

import (
	"encoding/json"
	"errors"
	"github.com/patrickmn/go-cache"
	"strings"
	"sync"
	"time"
)

type BaseCache struct {
	Key string `json:"key"`
}

type SetCache struct {
	BaseCache
	Value  interface{} `json:"value"`
	Expire int64       `json:"expire"`
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

func cacheGet(arg string) (interface{}, error) {
	var bc BaseCache
	err := json.Unmarshal([]byte(arg), &bc)
	if err != nil {
		return "", err
	}
	foo, found := c.Get(bc.Key)
	if found {
		return foo, nil
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

func cacheIncr(arg string) (interface{}, error) {
	var ic IntCache
	err := json.Unmarshal([]byte(arg), &ic)
	if err != nil {
		return 0, err
	}
	if ic.Key == "" {
		return 0, errors.New("Key is empty ")
	}
	mu := sync.RWMutex{}
	mu.Lock()
	err = c.Increment(ic.Key, ic.Step)
	if err != nil {
		mu.Unlock()
		c.Set(ic.Key, ic.Step, -1)
	}
	foo, found := c.Get(ic.Key)
	mu.Unlock()
	if found {
		return foo, nil
	}
	return 0, errors.New("Not Found ")
}

func cacheDecr(arg string) (int64, error) {
	var ic IntCache
	err := json.Unmarshal([]byte(arg), &ic)
	if err != nil {
		return 0, err
	}
	if ic.Key == "" {
		return 0, errors.New("Key is empty ")
	}
	err = c.Decrement(ic.Key, ic.Step)
	if err != nil {
		return 0, err
	}
	foo, found := c.Get(ic.Key)
	if found {
		return foo.(int64), nil
	}
	return 0, errors.New("Not Found ")
}
