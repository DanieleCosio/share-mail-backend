package url

import (
	"context"
	"sharemail/internal/db"

	"github.com/redis/go-redis/v9"
)

func GetUrlsList(key string) (*[]string, error) {
	client, err := db.GetRedisConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	busy, err := client.LRange(ctx, key, 0, -1).Result()
	if err == redis.Nil {
		return &[]string{}, nil
	} else if err != nil {
		return nil, err
	}

	return &busy, nil

}
