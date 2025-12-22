package redisService

import (
	"context"
	"go-backend-api/internal/global"
	"log"
)

var ctx = context.Background()

func GetFromDetailKey(id any, key string) any {
	detailKey := DetailKey(id, key)

	var returnValue any

	returnValue, err := global.Rdb.Get(ctx, detailKey).Result()

	if err != nil {
		return nil
	}
	log.Print("Saved success")
	return returnValue
}

func SetToDetailKey(id any, key string, value any) {
	detailKey := DetailKey(id, key)

	err := global.Rdb.Set(ctx, detailKey, value, 0).Err()
	if err != nil {
		log.Print("Not save success")
		return
	}
	log.Print("Saved success")
}

func GetFromListKey(key string) any {
	listKey := ListKey(key)

	var returnValue any

	returnValue, err := global.Rdb.Get(ctx, listKey).Result()

	if err != nil {
		return nil
	}
	log.Print("Saved success")
	return returnValue
}

func SetToListKey(key string, value any) {
	listKey := ListKey(key)

	err := global.Rdb.Set(ctx, listKey, value, 0).Err()
	if err != nil {
		log.Print("Not save success")
		return
	}
	log.Print("Saved success")
}

func DeleteKey(key string) {
	err := global.Rdb.Del(ctx, key).Err()
	if err != nil {
		log.Print("Not delete success")
		return
	}
	log.Print("Delete success")
}

func SetToUserTokenKey(token string, value any) {
	userTokenKey := UserTokenKey(token)

	err := global.Rdb.Set(ctx, userTokenKey, value, 0).Err()
	if err != nil {
		log.Print("Not save success")
		return
	}
	log.Print("Saved success")
}

func GetFromUserTokenKey(token string) any {
	userTokenKey := UserTokenKey(token)

	var returnValue any

	returnValue, err := global.Rdb.Get(ctx, userTokenKey).Result()

	if err != nil {
		return nil
	}
	log.Print("Saved success")
	return returnValue
}
