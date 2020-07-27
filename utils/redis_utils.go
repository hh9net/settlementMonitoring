package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"reflect"
)

/*
type RedisDB struct {
	RedisdbConfig *RedisDBConfig
	Lock          sync.RWMutex // lock
	RedisClient   *redis.Conn  // mysql client
}

type RedisDBConfig struct {
	RedisDBAddr   string
	redisPwd      string
	RedisDatabase uint8
}

var RedisConn *redis.Conn

func NewRedis() {
	RedisConn = Redisinit(&RedisDBConfig{
		RedisDBAddr: config.Optional.RedisAddr,
		redisPwd:config.Optional.RedisPwd,
		RedisDatabase: config.Optional.RedisDB,
	})
}
*/
//redis 工具箱
//1、redis 初始化
func RedisInit() *redis.Conn {
	logrus.Infoln("starting redis")
	//连接数据库
	//address:="192.168.200.170:6379"
	address := "127.0.0.1:6379"
	//address :=redisconf.RedisDBAddr
	conn, err := redis.Dial("tcp", address /*,redis.DialPassword("123456")*/)
	if err != nil {
		panic(err)
	}
	fmt.Println(address, "redis连接成功!")
	return &conn
}

//连接redis数据库
func RedisSelectDB(conn *redis.Conn /*redisconf *RedisDBConfig*/) {
	reply, serr := (*conn).Do("select", 3)
	if serr != nil {
		panic(serr)
	}
	fmt.Println("select 3 reply:", reply) //设置成功: ok
}

//set设置值
func RedisSet(conn *redis.Conn, key string, value string) {
	RedisSelectDB(conn)
	result, err := (*conn).Do("SET", key, value)
	if err != nil {
		panic(err)
	}
	fmt.Println("set result:", result) //设置成功，ok
}

//get value
func RedisGet(conn *redis.Conn, key string) {
	RedisSelectDB(conn)
	value, err := (*conn).Do("GET", key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get value: v=%s\n", value) //get value
}

//hset 设置值
func RedisHSet(conn *redis.Conn, key string, item string, value string) {
	RedisSelectDB(conn)
	//hset
	_, err := (*conn).Do("HSet", key, item, value)
	if err != nil {
		fmt.Println("hset出错，错误信息：", err)
		return
	}
	fmt.Println("hset ok")
}

//hget设置值
func RedisHGet(conn *redis.Conn, key string, item string) {
	RedisSelectDB(conn)
	value, err := (*conn).Do("HGET", key, item)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get value: value=%s\n", value) //get value
}

func RedisExample() {
	//连接数据库
	//address:="192.168.200.170:6379"
	address := "127.0.0.1:6379"
	conn, err := redis.Dial("tcp", address /*,redis.DialPassword("123456")*/)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("redis连接成功!")

	////测试是否连接到了redis
	//result, err := conn.Do("PING")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)

	reply, serr := conn.Do("select", 3)
	if serr != nil {
		panic(serr)
	}
	fmt.Println("select 3 reply:", reply) //设置成功: ok
	//1、设置key
	result, err := conn.Do("SET", "keyabc123456", "abc123")
	if err != nil {
		panic(err)
	}
	fmt.Println("set result:", result) //设置成功，ok

	//查询key的类型
	/*result, err = conn.Do("TYPE","key123")
	if err != nil {
		panic(err)
	}
	fmt.Println(result,reflect.TypeOf(result))//当返回none 表示key不存在*/

	//判断key是否存在
	/*	result, err = conn.Do("EXISTS", "key")
		if err != nil {
			panic(err)
		}
		fmt.Println(result, reflect.TypeOf(result)) //1表示存在，0表示不存在*/

	//2、设置过期时间 expire设置过期秒数  expireat设置以Unix时间戳的方式过期
	result, err = conn.Do("EXPIRE", "keyabc123456", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(result, reflect.TypeOf(result))

	//删除key
	/*result, err = conn.Do("DEL", "key")
	if err != nil {
		panic(err)
	}
	fmt.Println(result, reflect.TypeOf(result))*/

	//2.通过go向redis写入数据 string [key-value]
	//hset
	_, err = conn.Do("HSet", "user1", "name", "Bob小米")
	if err != nil {
		fmt.Println("hset出错，错误信息：", err)
		return
	}
	_, err = conn.Do("HSet", "user1", "age", "18")
	if err != nil {
		fmt.Println("hset出错，错误信息：", err)
		return
	}
	fmt.Println("hset ok")

	//3.通过go向redis读取数据 string [key-value]
	//因为返回的reply是{}interface。而name对应的值是string，因此需要转换
	reply1, err := redis.String(conn.Do("HGet", "user1", "name"))
	if err != nil {
		fmt.Println("hget出错，错误信息：", err)
		return
	}
	reply2, err := redis.Int(conn.Do("HGet", "user1", "age"))
	if err != nil {
		fmt.Println("hget出错，错误信息：", err)
		return
	}
	fmt.Println("hget ok")
	fmt.Printf("名字=%v 年龄=%v\n", reply1, reply2)

	//HMSet/HMGet(批量，一次性hset/hget多个field-value数据)

	//2.通过go向redis写入数据
	_, err = conn.Do("HMSet", "user2", "name", "Bob大米", "age", "19")
	if err != nil {
		fmt.Println("hmset出错，错误信息：", err)
		return
	}

	fmt.Println("hmset ok")

	//3.通过go向redis读取数据
	replys, hmgeterr := redis.Strings(conn.Do("HMGet", "user2", "name", "age"))
	if hmgeterr != nil {
		fmt.Println("hmget出错，错误信息：", hmgeterr)
		return
	}
	fmt.Println("hmget ok")

	for i, v := range replys {
		fmt.Printf("[%d]=%s\t", i, v)
	}
}

/*// 当只连接一个数据源的时候，可以直接使用GormClient
// 否则应当自己持有管理InitGormDB返回的GormDB
var RedisClient *RedisDB

type RedisDB struct {
	redisConfig *RedisConfig
	Client      *redis.Client
	lock        sync.RWMutex // lock
}

type RedisConfig struct {
	RedisAddr string
	RedisPwd  string
	RedisDB   int
}

func InitRedis(redisConfig *RedisConfig) *RedisDB {
	redisClient := &RedisDB{
		redisConfig: redisConfig,
		lock:        sync.RWMutex{},
		Client: redis.NewClient(&redis.Options{
			Addr:     redisConfig.RedisAddr,
			Password: redisConfig.RedisPwd, // no password set
			DB:       redisConfig.RedisDB,  // use default DB
		}),
	}
	_, err := redisClient.Client.Ping().Result()
	if err != nil {
		logrus.WithField("redisConfig", redisConfig).Errorln("ping redis error!")
	}
	go redisClient.redisTimer(redisConfig)
	RedisClient = redisClient
	return redisClient
}

func (p *RedisDB) reconnect() {
	client := redis.NewClient(&redis.Options{
		Addr:     p.redisConfig.RedisAddr,
		Password: p.redisConfig.RedisPwd, // no password set
		DB:       p.redisConfig.RedisDB,  // use default DB
	})
	p.Client = client
	_, err := p.Client.Ping().Result()
	if err != nil {
		logrus.WithField("redisConfig", p.redisConfig).Errorln("ping redis error!")
	}
}

func (p *RedisDB) redisTimer(redisConfig *RedisConfig) {
	redisTicker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-redisTicker.C:
			_, err := p.Client.Ping().Result()
			if err != nil {
				logrus.Errorln("redis connect fail,err:", err)
				p.reconnect()
			}
		}
	}
}
*/
