package redis

import (
	"github.com/streamcord/commands/internal"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var (
	// Context - the context passed to Redis when performing operations
	Context = context.Background()
	// RedisClient - the client connected to Redis
	RedisClient *redis.Client
)

// ConnectToRedis - connects to the Redis server
func ConnectToRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     internal.GlobalConf.Redis.Host,
		Password: internal.GlobalConf.Redis.Password,
		DB:       internal.GlobalConf.Redis.Database,
	})

	log.Info("Successfully connected to Redis")
}
