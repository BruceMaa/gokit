package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedigoConfig struct {
	Rawurl         string        `json:"rawurl"` // redis://:password@hostname:port/db_number
	Hostname       string        `json:"hostname"`
	Port           int64         `json:"port"`
	DbNumber       int           `json:"db_number"`
	Password       string        `json:"password,omitempty"`
	ConnectTimeout time.Duration `json:"connect_timeout"`
	KeepAlive      time.Duration `json:"keep_alive"`
	ReadTimeout    time.Duration `json:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout"`
}

type RedigoCli struct {
	redis.Conn
}

// 新建一个redigo客户端连接redis，不使用TLS
func (rc *RedigoConfig) NewConnection() (redigoCli RedigoCli, err error) {
	var options = make([]redis.DialOption, 0)
	if rc.ConnectTimeout != 0 {
		options = append(options, redis.DialConnectTimeout(rc.ConnectTimeout))
	}
	if rc.KeepAlive != 0 {
		options = append(options, redis.DialKeepAlive(rc.KeepAlive))
	}
	if rc.ReadTimeout != 0 {
		options = append(options, redis.DialReadTimeout(rc.ReadTimeout))
	}
	if rc.WriteTimeout != 0 {
		options = append(options, redis.DialWriteTimeout(rc.WriteTimeout))
	}
	if rc.Rawurl != "" {
		redigoCli.Conn, err = redis.DialURL(rc.Rawurl, options...)
	} else {
		redisDB := redis.DialDatabase(rc.DbNumber)
		redisPwd := redis.DialPassword(rc.Password)
		options = append(options, redisDB, redisPwd)
		redigoCli.Conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", rc.Hostname, rc.Port), options...)
	}

	pong, err := redis.String(redigoCli.Do("PING"))
	if pong != "PONG" {
		err = fmt.Errorf("redis connect fail\n")
	}
	return
}

/**************************String start***********************/

// 取值
func (rc *RedigoCli) GetString(key string) (string, error) {
	return redis.String(rc.Do("GET", key))
}

// range取值
func (rc *RedigoCli) GetRangeString(key string, start, end int) (string, error) {
	return redis.String(rc.Do("GETRANGE", key, start, end))
}

// 取值同时赋值
func (rc *RedigoCli) GetSetString(key, value string) (string, error) {
	return redis.String(rc.Do("GETSET", key, value))
}

// int64自增加1
func (rc *RedigoCli) IncrString(key string) (int64, error) {
	return redis.Int64(rc.Do("INCR", key))
}

// int64增加
func (rc *RedigoCli) IncrbyString(key string, increment int64) (int64, error) {
	return redis.Int64(rc.Do("INCRBY", key, increment))
}

// float64增加
func (rc *RedigoCli) IncrbyfloatString(key string, increment float64) (float64, error) {
	return redis.Float64(rc.Do("INCRBYFLOAT", key, increment))
}

// 批量取值
func (rc *RedigoCli) MGetString(keys ...interface{}) ([]string, error) {
	return redis.Strings(rc.Do("MGET", keys...))
}

// 批量赋值
func (rc *RedigoCli) MSetString(items map[string]string) (string, error) {
	var options = make([]interface{}, 0)
	for key, value := range items {
		fmt.Println("key:", key, " value:", value)
		options = append(options, key, value)
	}
	fmt.Println(options)
	return redis.String(rc.Do("MSET", options...))
}

// 赋值
func (rc *RedigoCli) SetString(key, value string) (string, error) {
	return redis.String(rc.Do("SET", key, value))
}

// 临时赋值,时间：秒
func (rc *RedigoCli) SetStringTemp(key, value string, second int) (string, error) {
	return redis.String(rc.Do("SET", key, value, "EX", second))
}

/**************************String end***********************/
