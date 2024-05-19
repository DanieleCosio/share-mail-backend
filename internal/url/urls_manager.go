package url

import (
	"context"
	"sharemail/internal/config"
	"sharemail/internal/db"
	"slices"

	"github.com/redis/go-redis/v9"
)

func mapDiff(m1, m2 *[]string) *[]string {
	var diff = []string{}
	for _, v := range *m1 {
		if !slices.Contains(*m2, v) {
			diff = append(diff, v)
		}
	}

	return &diff
}

func SyncUrls(urls *[]string) error {
	busy, err := GetUrlsList(config.AppConfig["BUSY_REDIS_KEY"])
	if err != nil {
		return err
	}

	free, err := GetUrlsList(config.AppConfig["FREE_REDIS_KEY"])
	if err != nil {
		return err
	}

	busyDiff := mapDiff(urls, busy)
	freeDiff := mapDiff(busyDiff, free)

	ctx := context.Background()
	redisClient, err := db.GetRedisConnection()
	if err != nil {
		return err
	}

	err = redisClient.Watch(ctx, func(tx *redis.Tx) error {
		exist, err := tx.Exists(ctx, config.AppConfig["FREE_REDIS_KEY"]).Result()
		if err != nil {
			return err
		}

		if exist == 1 {
			tx.Del(ctx, config.AppConfig["FREE_REDIS_KEY"])
		}

		_, err = tx.RPush(ctx, config.AppConfig["FREE_REDIS_KEY"], (*freeDiff)[0]).Result()
		if err != nil {
			return err
		}

		if len((*freeDiff)) <= 1 {
			return nil
		}

		_, err = tx.RPushX(ctx, config.AppConfig["FREE_REDIS_KEY"], (*freeDiff)[1:]).Result()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
