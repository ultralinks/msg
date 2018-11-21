package app

import (
	"fmt"

	"github.com/go-redis/redis"
)

var CodeCache *redis.Client

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func InitCache() {
	CodeCache = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Addr,
		Password: Config.Redis.Password,
		DB:       Config.Redis.DB, // use default DB
	})
	pong, err := CodeCache.Ping().Result()
	fmt.Println("pong", pong)
	fmt.Println("err", err)

}