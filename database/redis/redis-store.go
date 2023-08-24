package redis

import (
	"fmt"
	"time"
)

const CAPTCHA = "captcha:"

type Store struct{}

// Set 实现设置captcha的方法
func (Store) Set(id string, value string) error {
	key := CAPTCHA + id
	err := rdb.Set(ctx, key, value, time.Minute*2).Err()
	return err
}

// Get 实现获取captcha的方法
func (Store) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := rdb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify 实现验证captcha的方法
func (Store) Verify(id, answer string, clear bool) bool {
	v := Store{}.Get(id, clear)
	return v == answer
}
