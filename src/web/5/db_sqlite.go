package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	checkErr(err)

	stmt,err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec("lyt", "研发", "2016-06-03")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("lyt_update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)


	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid, username, department, created)
	}


	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)


	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
