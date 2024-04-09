package dao

import (
	"log"
	"newRpc/model"
)

func GetUser() ([]model.User, error) { //获取用户信息
	var users []model.User //定义切片存储用户
	s := "SELECT ID, Username, Password FROM users "
	rows, err := DB.Query(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Username, &u.Password) //查询数据
		if err != nil {
			log.Println(err)
			return nil, err
		}

		users = append(users, u) //导入切片中
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func VerifyUserAndGetID(username, password string) (int64, error) { //确认用户id避免出现同名
	var userID int64
	s := "SELECT ID FROM user WHERE Username=? AND Password=?" //检索出用户的id
	err := DB.QueryRow(s, username, password).Scan(&userID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return userID, nil
}

func AddUser(u model.User) error {
	s := "INSERT INTO user (username, password) VALUES (?, ?)" //插入注册的信息
	result, err := DB.Exec(s, u.Username, u.Password)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return err

	}

	u.ID = id
	return nil
}
