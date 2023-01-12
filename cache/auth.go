package cache

import (
	"context"
	"fmt"
)

const (
	KeyPrefix_UserSessionKey = "sessionkey::userid::%s"
)

func GetUserSession(ctx context.Context, userid string) (sessionKey string, err error) {
	return redisClient.Get(ctx, userid).Result()
}

func SetUserSession(ctx context.Context, userid string, sessionKey string) (string, error) {
	key := fmt.Sprintf(KeyPrefix_UserSessionKey, userid)
	return redisClient.Set(ctx, key, sessionKey, 0).Result()
}
