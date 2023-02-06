/*
 * @Description: sql使用
 */
package model

// database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动

import (
	"database/sql"
	"fmt"
	"project/conf"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
)

func Datebase() {
	// 运行加载ini参数
	conf.LoadMysqlData()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.DbUser, conf.DbPassWord, conf.DbHost, conf.DbPort, conf.DbName)
	// Open函数可能只是验证其参数格式是否正确，
	// 实际上并不创建与数据库的连接
	fmt.Println(dsn)
		db, err := sql.Open(conf.Db, dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)  //设置打开最大连接
	db.SetMaxIdleConns(20)   //设置连接池 空闲
	db.SetConnMaxLifetime(time.Second * 30)
	Db = db
}
