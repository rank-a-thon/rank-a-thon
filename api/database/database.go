package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_redis "github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func Init() {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	var err error
	db, err = ConnectDB(dbInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func InitForTest() {
	dbInfo := "host=localhost port=5432 user=postgres dbname=rankathon_test sslmode=disable"
	fmt.Println(dbInfo)
	var err error
	db, err = ConnectDB(dbInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}

var RedisClient *_redis.Client

func InitRedis(params ...string) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	db, _ := strconv.Atoi(params[0])

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       db,
	})
}

func GetRedis() *_redis.Client {
	return RedisClient
}
