package mysql

import (
	"log"
	"testing"
)

func TestMySQLConfig_NewConnection(t *testing.T) {
	mysqlConfig := &MySQLConfig{
		Rawurl: "rq8hxb1ru3:0A3r059F7FQd@tcp(rm-bp16gxb8lj09v0u4j.mysql.rds.aliyuncs.com:3306)/rq8hxb1ru3",
	}

	mysqlCli, err := mysqlConfig.NewConnection()
	if err != nil {
		log.Fatalln(err)
	}
	defer mysqlCli.Close()

	log.Println(mysqlCli.Ping())
	log.Println(mysqlCli.Stats())
}
