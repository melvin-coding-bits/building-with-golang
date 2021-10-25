package redis

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/melvinodsa/build-with-golang/api-gateway/config"
)

func SaveUserInfoToRedis(appCtx *config.AppContext, userId, role, authToken string) error {
	hashedAuthToken := sha256.Sum256([]byte(authToken))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return appCtx.Redis.SetEX(ctx, hex.EncodeToString(hashedAuthToken[:]), userId+"|"+role, time.Hour*24*7).Err()
}

func GetUserInfoFromRedis(appCtx *config.AppContext, authToken string) (string, string, error) {
	hashedAuthToken := sha256.Sum256([]byte(authToken))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response := appCtx.Redis.Get(ctx, hex.EncodeToString(hashedAuthToken[:]))
	if response.Err() != nil {
		return "", "", response.Err()
	}
	value := response.Val()
	values := strings.Split(value, "|")
	if len(values) != 2 {
		return "", "", errors.New("invalid user info")
	}

	return values[0], values[1], nil
}
