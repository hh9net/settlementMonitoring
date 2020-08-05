package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
)

//golang操作oracle数据库

//执行SQL语句
func sqlExec(db *sql.DB, sqlStmt string) error {
	res, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("SQL Execute success rows affected %d\n", num)
	return nil
}

//sql查询
func sqlQuery(db *sql.DB, sqlStmt string) error {
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

	fmt.Print("连接数据库")
	//fmt.Print(os.Environ())
	//db, err := sql.Open("oci8", "admin/123@//192.168.0.160:1521/ORCL")

	//db, err := sql.Open("oci8", fmt.Sprintf("%s/%s@%s", "system", "oracle", "127.0.0.1:49166/orcl"))//docker 端口映射 49166：1521

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := sqlExec(db, "create table mytest(name varchar2(10), age int, primary key(name))"); err != nil {
		log.Fatal(err)
	}

	if err := sqlExec(db, "insert into mytest(name, age) values('Tom', 20)"); err != nil {
		log.Fatal(err)
	}

	if err := sqlExec(db, "insert into mytest(name, age) values('Jerry', 20)"); err != nil {
		log.Fatal(err)
	}
	if err := sqlQuery(db, "SELECT * FROM (SELECT t.*  FROM ADMIN.B_TXF_CHEDXFYSSJ t    ) WHERE ROWNUM <= 5"); err != nil {
		log.Fatal(err)
	}

	if err := sqlExec(db, "drop table mytest"); err != nil {
		log.Fatal(err)
	}
}
