package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("hello world")


	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.90:36379",
		Password: "XqQjF&80gXZANz", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
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
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}
