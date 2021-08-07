package dbClient

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jakecallery/iiria/apiServer/data"
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

func (c *RedisClient) DataFromTime(t string) *data.WeatherData {
	c.l.Println("data from time")
	return &data.WeatherData{}
}

func (c *RedisClient) CheckConnection() error {
	c.l.Println("Checking Connection...")
	_, err := c.redisClient.Ping(c.ctx).Result()

	if err == nil {
		c.l.Println("Connection is good!")
	}

	return err
}
