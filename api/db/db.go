package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

var db *gorp.DbMap

func Init() {

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

func GetDB() *gorp.DbMap {
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
