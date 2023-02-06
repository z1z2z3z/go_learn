/*
 * @Description: 创建数据库
 */
package main

import (
	"fmt"
	"project/model"
	_ "github.com/go-sql-driver/mysql"
)

type vedio struct {
	id         int
	vedio_name string
	vedio_id   int
}

func main() {
	model.Datebase()
	prepareInsertDemo()
	queryRowDemo()
	queryMultiRowDemo()

}

// 查询单条数据实例
func queryRowDemo() {
	var v vedio
	sqlStr := "select * from table_colly where id = ?"
	// QueryRow之后需要调用Scan方法，否则持有的数据库链接不会被释放
	err := model.Db.QueryRow(sqlStr, 1).Scan(&v.id, &v.vedio_name, &v.vedio_id)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("one vedio:", v)
}


// 预处理查询多条数据实例
func queryMultiRowDemo()  {
	sqlStr := "select * from table_colly where id > ?"
	stmt, err := model.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var ve vedio
		err := rows.Scan(&ve.id, &ve.vedio_name, &ve.vedio_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Println("each vedio:",ve)
	}
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into table_colly(vedio_name, vedio_id) values (?,?)"
	stmt, err := model.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18454545)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("肖申克的救赎", 14545142)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}