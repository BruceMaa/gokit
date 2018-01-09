package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

//func TestMySQLConfig_NewConnection(t *testing.T) {
//	mysqlConfig := &MySQLConfig{
//		Rawurl:          "rq8hxb1ru3:0A3r059F7FQd@tcp(rm-bp16gxb8lj09v0u4j.mysql.rds.aliyuncs.com:3306)/rq8hxb1ru3",
//		MaxOpenConns:    10,
//		MaxIdleConns:    10,
//		ConnMaxLifetime: 5 * time.Minute,
//	}
//
//	mysqlCli, err := mysqlConfig.NewConnection()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer mysqlCli.Close()
//
//	//log.Println(mysqlCli.Ping())
//	//log.Println(mysqlCli.Stats())
//
//	//args := [][]interface{}{{"test1", "test1"}, {"test2", "test2"}, {"test3", "test3"}, {"test4", "test4"}}
//	//_, err = mysqlCli.BatchInsert("insert person(first_name, last_name) values(?, ?)", args)
//
//	r, err := mysqlCli.Query("select first_name, last_name from person where id = ?", 2)
//	var first_name, last_name string
//	r.Scan(&first_name, &last_name)
//	fmt.Println(first_name, last_name)
//	fmt.Println(err)
//}

func TestMySQLCli_Query(t *testing.T) {
	db, err := sql.Open("mysql", "rq8hxb1ru3:0A3r059F7FQd@tcp(rm-bp16gxb8lj09v0u4j.mysql.rds.aliyuncs.com:3306)/rq8hxb1ru3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("cols:", cols)
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))

	for i := range vals {
		scans[i] = &vals[i]
	}

	fmt.Println("scans:", scans)

	var results []map[string]string

	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			log.Fatalln(err)
		}

		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}

	for k, v := range results {
		fmt.Println(k, v)
	}
}
