package cache

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"newRpc/dao"
	"newRpc/model"
	"time"
)

func UpdateRedisRank(ctx context.Context, client *redis.Client, db *sql.DB) error {

	statsLIst, err := dao.GetUrlVisits(dao.DB) //获取点击量
	if err != nil {
		log.Println(err)
		return err
	}

	for i, stats := range statsLIst {
		stats.Rank = i + 1
		statsJSON, err := json.Marshal(stats)
		if err != nil {
			log.Println(err)
			return err
		} //进行排名

		err = client.Set(ctx, stats.ShortUrl+"_stats", statsJSON, time.Hour).Err()
		if err != nil {
			log.Println(err)
			return err
		} //缓存存储排名

	}

	return nil
}

func GetAllRank(ctx context.Context, client *redis.Client) (map[string]int, error) {
	//更新排名
	err := UpdateRedisRank(ctx, client, dao.DB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 获取所有的短链接排名数据
	keys, err := client.Keys(ctx, "*_stats").Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 存储所有短链接的排名
	ranks := make(map[string]int)
	for _, key := range keys {
		// 从 Redis 中获取排名数据的 JSON 字符串
		statsJSON, err := client.Get(ctx, key).Result()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// 解析 JSON 字符串
		var stats model.URLStats
		if err := json.Unmarshal([]byte(statsJSON), &stats); err != nil {
			log.Println(err)
			return nil, err
		}

		// 存储排名字段
		ranks[stats.ShortUrl] = stats.Rank
	}

	return ranks, nil
}
