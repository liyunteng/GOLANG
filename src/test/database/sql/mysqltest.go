package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

// SQL
// CREATE TABLE IF NOT EXISTS `test`.`user` (
//  `user_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
//  `user_name` VARCHAR(45) NOT NULL COMMENT '用户名称',
//  `user_age` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄',
//  `user_sex` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
//  PRIMARY KEY (`user_id`))
//  ENGINE = InnoDB
//  AUTO_INCREMENT = 1
//  DEFAULT CHARACTER SET = utf8
//  COLLATE = utf8_general_ci
//  COMMENT = '用户表'

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// insert
func insert(db *sql.DB) {
	stmt, err := db.Prepare(`INSERT user (user_name, user_age, user_sex) values (?, ?, ?)`)
	checkErr(err)

	res, err := stmt.Exec("tony", 20, 1)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func query(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	// for rows.Next() {
	//	var userId int
	//	var userName string
	//	var userAge int
	//	var userSex int

	//	rows.Columns()
	//	err = rows.Scan(&userId, &userName, &userAge, &userSex)
	//	checkErr(err)

	//	fmt.Println(userId)
	//	fmt.Println(userName)
	//	fmt.Println(userAge)
	//	fmt.Println(userSex)
	// }

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}

		fmt.Printf("[id:%v\tname:%v\tage:%v\tsex:%v]\n",
			record["user_id"], record["user_name"],
			record["user_age"], record["user_sex"]);
	}

	fmt.Println()
}

func update(db *sql.DB)  {
	stmt, err := db.Prepare(`UPDATE user set user_age=?,user_sex=? WHERE user_name=?`)
	checkErr(err)

	res, err := stmt.Exec(21, 2, "tony")
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(num)
}

func remove(db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM user WHERE user_name=?`)
	checkErr(err)

	res, err := stmt.Exec("tony")
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(num)
}



func main() {
	db, err := sql.Open("mysql", "lyt@/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	query(db)
	insert(db)
	query(db)
	update(db)
	query(db)
	remove(db)
	query(db)
}
