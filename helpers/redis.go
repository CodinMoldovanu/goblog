package helpers

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func getRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.88.226:6379",
		Password: "",
		DB:       0,
	})

	return client
}

// SetUserToken into Redis database for easy retrieval
func SetUserToken(token string, userid int) {
	client := getRedisClient()
	defer client.Close()
	duration := time.Hour * 72

	client.Set(strconv.Itoa(userid), token, duration)
}

// GetUserToken and check if same
func GetUserToken(userid int) *redis.StringCmd {
	client := getRedisClient()
	defer client.Close()
	token := client.Get(strconv.Itoa(userid))

	return token
}
