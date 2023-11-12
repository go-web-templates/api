package data

import (
	"fmt"

	"github.com/go-web-templates/api/pkg/conf"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Ctx *redis.Client
}

func NewCache(appConf *conf.AppConf) *Cache {
	cacheConf := appConf.Data.Cache

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cacheConf.Host, cacheConf.Port),
		Password: cacheConf.Pass,
		DB: 0,
	})

	return &Cache{ Ctx: redisClient }
}
