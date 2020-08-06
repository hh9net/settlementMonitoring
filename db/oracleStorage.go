package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
	"os"
)

//golang操作oracle数据库
//执行SQL语句
func sqlExec(db *sql.DB, sqlStmt string) (error, int64) {
	res, err := db.Exec(sqlStmt)
	if err != nil {
		return err, 0
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err, 0
	}
	log.Printf("SQL Execute success rows affected %d\n", num)
	return nil, num
}

//sql查询
func sqlQuery(db *sql.DB, sqlStmt string) (error, int) {
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return err, 0
	}
	defer rows.Close()
	var n int
	for rows.Next() {
		n++
		//log.Printf("row[%d]\n", n )
	}
	err = rows.Err()
	if err != nil {
		return err, 0
	}
	log.Printf("SQL Query success rows queried %d\n", n)
	return nil, n
}

//查询海玲同步结算数量 B_TXF_CHEDXFYSSJ
func OrclQuerydata() int {
	// 为log添加短文件名,方便查看行数
	log.Println("Oracle Driver example")
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	db, err := sql.Open("oci8", fmt.Sprintf("%s/%s@%s", "admin", "123", "192.168.0.160:1521/orcl"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接数据库成功", db)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var num int ///*  where   F_VC_JIAOYZT   = 99*/ CHEDXFYSSJ
	if err, num = sqlQuery(db, "  SELECT  F_VC_JIAOYZT  FROM  B_TXF_CHEDXFYSSJ    "); err != nil {
		log.Fatal(err)
	}
	return num
}

//测试
//sql查询
func testsqlQuery(db *sql.DB, sqlStmt string) error {
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	var n int
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			return err
		}
		n++
		log.Printf("row[%d], name=[%s], age=[%d]\n", n, name, age)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	log.Printf("SQL Query success rows queried %d\n", n)
	return nil
}
func Orcldb() {
	// 用户名/密码@IP:端口/实例名  admin/123@192.168.0.160:1521/orcl
	db, err := sql.Open("oci8", fmt.Sprintf("%s/%s@%s", "admin", "123", "192.168.0.160:1521/orcl"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接数据库成功")
	//fmt.Print(os.Environ())
	//db, err := sql.Open("oci8", "admin/123@//192.168.0.160:1521/ORCL")

	//db, err := sql.Open("oci8", fmt.Sprintf("%s/%s@%s", "system", "oracle", "127.0.0.1:49166/orcl"))//docker 端口映射 49166：1521

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err, _ := sqlExec(db, "create table mytest(name varchar2(10), age int, primary key(name))"); err != nil {
		log.Fatal(err)
	}

	if err, _ := sqlExec(db, "insert into mytest(name, age) values('Tom1111', 20)"); err != nil {
		log.Fatal(err)
	}

	if err, _ := sqlExec(db, "insert into mytest(name, age) values('Jerry1111', 20)"); err != nil {
		log.Fatal(err)
	}
	if err := testsqlQuery(db, "SELECT   *   FROM  MYTEST   "); err != nil {
		fmt.Println("SELECT error")
		log.Fatal(err)
	}
	if err, _ := sqlExec(db, "drop table mytest"); err != nil {
		log.Fatal(err)
	}
}
