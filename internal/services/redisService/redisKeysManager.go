package redisService

func DetailKey(id any, key string) string {
	return key + "-detail:" + id.(string)
}

func ListKey(key string) string {
	return key + "-list"
}
func UserTokenKey(token string) string {
	return "user-token:" + token
}
