package domain

const RedisUserEmail = "UserEmail"

func GetRedisKeyUserEmailVerify(email string) string {
	return RedisUserEmail + ":" + email + ":" + "verify"
}
