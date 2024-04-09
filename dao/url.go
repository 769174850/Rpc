package dao

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"newRpc/model"
	"time"
)

func GenerateUrl(LongUrl string) (shortUrl string, err error) {
	if LongUrl == "" {
		return "", err
	} // 检查URL是否为空

	// 生成短链接
	Hash := md5.New()
	Hash.Write([]byte(LongUrl))
	H := hex.EncodeToString(Hash.Sum(nil))

	shortUrl = H[:10]
	return shortUrl, nil
}

func InsertUrl(ShortUrl string, LongUrl string) error {

	s := "INSERT INTO short_urls (short_url, long_url, visits, created_at, update_at) VALUES (?, ?, ?, ?, ?)" //将数据插入数据库
	result, err := DB.Exec(s, ShortUrl, LongUrl, 0, time.Now(), time.Now())                                   //插入数据
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = result.LastInsertId() //检查是否插入成功
	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}

func DeleteUrl(id int64) error {
	s := "UPDATE short_urls SET is_deleted = ? where id = ?"

	result, err := DB.Exec(s, true, id) //删除数据
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = result.RowsAffected() //检查是否删除成功
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateUrl(oldShortUrl string) error {
	s := "UPDATE short_urls SET short_url = ?,update_at = ? where short_url = ?"

	newShortUrl, err := GenerateUrl(oldShortUrl) //使用旧的短链生成新的短链
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := DB.Exec(s, newShortUrl, time.Now(), oldShortUrl) //更新数据
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected() //检查是否更新成功
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(err)
		return err
	}
	return nil
}

func GetUserUrls(userId int64) ([]model.Url, error) {
	var Urls []model.Url
	s := "SELECT id, short_url, long_url, visits, user_id, created_at, updated_at FROM short_urls WHERE user_id = ? AND is_deleted = ?"
	rows, err := DB.Query(s, userId, false)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() { //遍历查询结果
		var url model.Url
		err := rows.Scan(&url.ID, &url.ShortUrl, &url.LongUrl, &url.Visits, &url.UserID, &url.CreatedAt, &url.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		Urls = append(Urls, url)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return Urls, nil
}

type URLStatsList []model.URLStats

func (usl URLStatsList) Len() int           { return len(usl) }
func (usl URLStatsList) Swap(i, j int)      { usl[i], usl[j] = usl[j], usl[i] }
func (usl URLStatsList) Less(i, j int) bool { return usl[i].Visits > usl[j].Visits } //对数据进行排行

func GetUrlVisits(db *sql.DB) ([]model.URLStats, error) {
	var statsList URLStatsList
	s := "SELECT id, short_url, visits FROM short_urls"

	rows, err := DB.Query(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close() //查询数据

	for rows.Next() {
		var id int64
		var visits, rank int
		var shortUrl string

		if err := rows.Scan(&id, &shortUrl, &visits); err != nil {
			log.Println(err)
			return nil, err
		}

		// 创建URLStats对象
		stats := model.URLStats{
			ID:       id,
			ShortUrl: shortUrl,
			Visits:   visits,
			Rank:     rank,
		}

		statsList = append(statsList, stats)
	}

	return statsList, err
}
