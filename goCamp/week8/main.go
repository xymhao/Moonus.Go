package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

func main() {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	client := redis.NewClient(options)
	defer client.Close()
	result, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	count := 1000

	group2 := sync.WaitGroup{}
	group2.Add(count)
	for i := 0; i < count; i++ {
		go Del(i, client, &group2)
	}
	group2.Wait()

	group := sync.WaitGroup{}
	group.Add(count)
	WriteTestKey(count, 10000, client, &group)
	group.Wait()

}

func Del(i int, client *redis.Client, w *sync.WaitGroup) {
	err := client.Del(fmt.Sprint(i)).Err()
	if err != nil {
		fmt.Println(err)
	}
	w.Done()
}

func WriteTestKey(count int, size int, client *redis.Client, group *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		var index = fmt.Sprint(i)
		bytes := make([]byte, size)
		err := client.Set(index, bytes, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		group.Done()
	}
}
