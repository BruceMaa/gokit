package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestRedigoConfig_NewConnection(t *testing.T) {
	redigoConfig := &RedigoConfig{
		//Rawurl:         "redis://:123456@pub-redis-15929.dal-05.1.sl.garantiadata.com:15929/0",
		Hostname:       "pub-redis-15929.dal-05.1.sl.garantiadata.com",
		Port:           15929,
		DbNumber:       0,
		Password:       "123456",
		ConnectTimeout: 5 * time.Minute,
		KeepAlive:      5 * time.Minute,
		WriteTimeout:   5 * time.Second,
		ReadTimeout:    5 * time.Second,
	}

	redisCli, err := redigoConfig.NewConnection()
	defer redisCli.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestRedigoCli_GetRangeString(t *testing.T) {
	redigoConfig := &RedigoConfig{
		Rawurl: "redis://:123456@pub-redis-15929.dal-05.1.sl.garantiadata.com:15929/0",
	}

	redisCli, err := redigoConfig.NewConnection()
	defer redisCli.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//r, _ := redisCli.GetRangeString("test", 6, 11)
	//r, _ := redisCli.GetSetString("test", "testnew")
	//r, err := redisCli.IncrString("mykey")
	//r, err := redisCli.IncrbyString("mykey", 4)
	//r, err := redisCli.IncrbyfloatString("mykey", 4.32)
	//r, err := redisCli.MGetString("mykey", "test")
	//r, err := redisCli.MSetString(map[string]string{"test1": "test1", "test2": "test2", "test3": "test3"})
	r, err := redisCli.SetStringTemp("test", "test1", 30)
	fmt.Println(r, err)
}
