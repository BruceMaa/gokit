package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySQLConfig struct {
	Rawurl          string        `json:"rawurl"`
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
}

type MySQLCli struct {
	*sql.DB
}

// 新建一个mysql连接，返回连接
func (mc *MySQLConfig) NewConnection() (mysqlCli MySQLCli, err error) {
	/*open函数并没有创建连接，它只是验证参数是否合法。然后开启一个单独goroutines去监听是否需要建立新的连接，当有请求建立新连接时就创建新连接。*/
	mysqlCli.DB, err = sql.Open("mysql", mc.Rawurl)
	if err != nil {
		return
	}

	if err = mysqlCli.Ping(); err != nil {
		return
	}
	mysqlCli.SetMaxOpenConns(mc.MaxOpenConns)
	mysqlCli.SetMaxIdleConns(mc.MaxIdleConns)
	mysqlCli.SetConnMaxLifetime(mc.ConnMaxLifetime)
	return
}
