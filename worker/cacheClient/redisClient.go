package cacheClient

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/jakecallery/iiria/worker/weatherClients"
)

type RedisClient struct {
	ServerAddr  string
	ServerPort  string
	ServerPass  string
	redisClient *redis.Client
	isReady     bool
	ctx         context.Context
	l           *log.Logger
}

func NewRedisClient(l *log.Logger) *RedisClient {
	c := RedisClient{
		ServerAddr: "localhost",
		ServerPort: "6379",
		ServerPass: "",
		isReady:    false,
		ctx:        context.Background(),
		l:          l,
	}

	return &c
}

func (c *RedisClient) Init() {
	c.redisClient = redis.NewClient(&redis.Options{
		Addr:     c.ServerAddr + ":" + c.ServerPort,
		Password: c.ServerPass,
		DB:       0,
	})

	c.isReady = true
}

//TODO: Implement RedisJSON
func (c *RedisClient) Save(wd *weatherClients.WeatherData) error {
	intervals := wd.Data.Timelines[0].Intervals

	for _, interval := range intervals {
		//TODO: Basic string santization/string checking
		st := strings.ReplaceAll(string(interval.StartTime), ":", "_")
		data, err := json.Marshal(interval.Values)
		if err != nil {
			c.l.Printf("Error marshaling json from weather to cache: %v", err)
			return err
		}
		c.redisClient.Set(c.ctx, st, string(data), 0)
	}

	return nil
}

func (c *RedisClient) CheckConnection() error {
	c.l.Println("Checking Connection...")
	_, err := c.redisClient.Ping(c.ctx).Result()

	if err == nil {
		c.l.Println("Connection is good!")
	}

	return err
}
