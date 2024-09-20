package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stathat/consistent"
)

const (
	dbUser     = "user"
	dbPassword = "password"
	dbHost     = "localhost"
	dbPort     = "3306"
)

var (
	consistentHash *consistent.Consistent
	dbPools        map[string]*sql.DB
)

func init() {
	// 初始化一致性哈希实例
	consistentHash = consistent.New()

	// 初始化数据库连接池
	dbPools = make(map[string]*sql.DB)
}

// 1、创建了连接池
func addDBPool(dbName string) {
	// 添加数据库实例到一致性哈希
	consistentHash.Add(dbName)

	// 创建数据库连接池并添加到 dbPools
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	dbPools[dbName] = db
}

// 2、获取数据库连接
// 创建一个函数来根据键（比如用户 ID）获取对应的数据库连接
func getDBConnectionByKey(key string) (*sql.DB, error) {
	dbName, err := consistentHash.Get(key)
	if err != nil {
		return nil, err
	}

	db, ok := dbPools[dbName]
	if !ok {
		return nil, fmt.Errorf("no db connection found for dbName: %s", dbName)
	}

	return db, nil
}

// 3.执行 SQL 语句
func getUserById(userId string) (*User, error) {
	db, err := getDBConnectionByKey(userId)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", userId).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type User struct {
}
