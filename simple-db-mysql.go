package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL数据库连接信息
	dbUser := "your_username"
	dbPass := "your_password"
	dbName := "your_database"
	dbHost := "localhost"
	dbPort := "3306"

	// 构建MySQL连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// 连接MySQL数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// 检测MySQL是否就绪
	err = db.Ping()
	if err != nil {
		log.Fatal("MySQL not ready:", err)
	}
	fmt.Println("Connected to MySQL!")

	// 执行增删改查操作的示例
	// 插入数据
	insertStatement := "INSERT INTO your_table (column1, column2) VALUES (?, ?)"
	_, err = db.Exec(insertStatement, "value1", "value2")
	if err != nil {
		log.Fatal("Insert error:", err)
	}
	fmt.Println("Data inserted successfully!")

	// 查询数据
	selectStatement := "SELECT * FROM your_table"
	rows, err := db.Query(selectStatement)
	if err != nil {
		log.Fatal("Select error:", err)
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var column1, column2 string
		err := rows.Scan(&column1, &column2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\n", column1, column2)
	}

	// 更新数据
	updateStatement := "UPDATE your_table SET column1 = ? WHERE column2 = ?"
	_, err = db.Exec(updateStatement, "new_value", "value2")
	if err != nil {
		log.Fatal("Update error:", err)
	}
	fmt.Println("Data updated successfully!")

	// 删除数据
	deleteStatement := "DELETE FROM your_table WHERE column1 = ?"
	_, err = db.Exec(deleteStatement, "new_value")
	if err != nil {
		log.Fatal("Delete error:", err)
	}
	fmt.Println("Data deleted successfully!")

	// 关闭数据库连接
	err = db.Close()
	if err != nil {
		log.Fatal("Error closing database connection:", err)
	}
}
