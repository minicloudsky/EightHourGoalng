package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLock struct {
	RedisClient *redis.Client
	// 热点key的锁，多个分布式锁的key名相同
	Key        string
	ExpireTime time.Duration
	// uuid，通常为当前时间的毫秒时间戳
	LockedToken string
}

func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return redisClient
}

func NewRedisLock(redisClient *redis.Client, key string, expireTime time.Duration) *RedisLock {
	return &RedisLock{
		RedisClient: redisClient,
		Key:         key,
		ExpireTime:  expireTime,
	}
}

func (lock *RedisLock) AcquireLock(ctx context.Context) (bool, error) {
	lock.LockedToken = fmt.Sprintf("%d", time.Now().UnixNano())
	return lock.RedisClient.SetNX(ctx, lock.Key, lock.LockedToken, lock.ExpireTime).Result()
}

func (lock *RedisLock) ReleaseLock() error {
	// 获取指定键的值，并与传入的参数进行比较。这样做的目的是为了确保当前脚本所持有的锁与传入的参数相匹配，
	// 即当前脚本所持有的锁仍然有效。如果不进行比较，而是直接删除指定的键，那么就无法确保当前脚本所持有的锁是有效的，
	// 可能会导致错误释放锁。因此，为了确保安全释放锁，需要先验证当前脚本所持有的锁与传入的参数是否匹配，
	// 只有在匹配的情况下才执行删除操作。
	luaScript := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`
	return lock.RedisClient.Eval(context.Background(), luaScript, []string{lock.Key}, lock.LockedToken).Err()
}

func QueryAndUpdateHotKey(bizName, hotKey string, newValue string) error {
	redisClient := NewRedisClient()
	lockKey := hotKey + ":lock"
	lock := NewRedisLock(redisClient, lockKey, 30*time.Second)
	acquired, err := lock.AcquireLock(context.Background())
	if err != nil {
		fmt.Printf("biz %s failed to get lock\n", bizName)
		return err
	}
	if acquired {
		defer func(lock *RedisLock) {
			err := lock.ReleaseLock()
			if err != nil {
				fmt.Println(err)
			}
		}(lock)
		// 模拟更新热点 key 的操作
		fmt.Printf("%s Updating hot key: %s\n", bizName, hotKey)
		// 更新热点 key
		err := redisClient.Set(context.Background(), hotKey, newValue, 30*time.Minute).Err()
		if err != nil {
			return err
		}
		fmt.Printf("%s Finish update hot key.\n", bizName)
	}
	return nil
}

func bizUpdateHotKey(bizName, hotKey string) {
	newValue := time.Now().String()
	err := QueryAndUpdateHotKey(bizName, hotKey, newValue)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	hotKey := "cloudsky_user_info"
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			bizUpdateHotKey(strconv.Itoa(i), hotKey)
		}(i)
	}
	wg.Wait()
}
