package utils

import (
	"testing"
)

//连接redis
func TestRedisInit(t *testing.T) {
	RedisInit()
}

//测试RedisSelectDB
func TestRedisSelectDB(t *testing.T) {
	RedisSelectDB(RedisInit())
}

func TestRedisExample(t *testing.T) {
	RedisExample()
}

//RedisSet
func TestRedisSet(t *testing.T) {
	RedisSet(RedisInit(), "kabc", "12345")
}

//RedisGet
func TestRedisGet(t *testing.T) {
	RedisGet(RedisInit(), "kabc")
}

//RedisHSet  覆盖set
func TestRedisHSet(t *testing.T) {
	//RedisHSet(RedisInit(), "tingabc12", "k1", "v1")
	RedisHSet(RedisInit(), "tingabc12", "k1", "v2")
}

//RedisHGet
func TestRedisHGet(t *testing.T) {
	RedisHGet(RedisInit(), "kabc12", "k1")
}
