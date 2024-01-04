package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Database struct represents the database connection
type Database struct {
	db *sql.DB
}

type UserData struct {
	Password                string
	LastLogin               sql.NullTime
	FirstName               string
	LastName                string
	IsActive                bool
	DateJoined              time.Time
	ID                      string
	Username                string
	Name                    string
	Email                   string
	Role                    string
	Avatar                  string
	Wechat                  string
	Phone                   sql.NullString
	PrivateKey              sql.NullString
	PublicKey               sql.NullString
	Comment                 sql.NullString
	IsFirstLogin            bool
	DateExpired             sql.NullString
	CreatedBy               sql.NullString
	MFALevel                int
	OTPSecretKey            sql.NullString
	Source                  string
	DatePasswordLastUpdated time.Time
	NeedUpdatePassword      bool
	DingtalkID              sql.NullString
	WecomID                 sql.NullString
	FeishuID                sql.NullString
	SecretKey               sql.NullString
	IsServiceAccount        bool
}

// NewDatabase creates a new Database instance
func NewDatabase(dbUser, dbPass, dbName, dbHost, dbPort string) (*Database, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// Close closes the database connection
func (d *Database) Close() error {
	return d.db.Close()
}

// InsertData inserts data into the specified table
func (d *Database) InsertData(table, column1, column2 string) error {
	insertStatement := fmt.Sprintf("INSERT INTO %s (column1, column2) VALUES ($1, $2)", table)
	_, err := d.db.Exec(insertStatement, column1, column2)
	return err
}

// QueryData queries data from the specified table
func (d *Database) QueryData(table string) (*sql.Rows, error) {
	selectStatement := fmt.Sprintf("SELECT * FROM %s", table)
	return d.db.Query(selectStatement)
}

// UpdateData updates data in the specified table
func (d *Database) UpdateData(table, newColumn1, oldColumn2 string) error {
	updateStatement := fmt.Sprintf("UPDATE %s SET column1 = $1 WHERE column2 = $2", table)
	_, err := d.db.Exec(updateStatement, newColumn1, oldColumn2)
	return err
}

// DeleteData deletes data from the specified table
func (d *Database) DeleteData(table, column1 string) error {
	deleteStatement := fmt.Sprintf("DELETE FROM %s WHERE column1 = $1", table)
	_, err := d.db.Exec(deleteStatement, column1)
	return err
}

func main() {
	// PostgreSQL数据库连接信息
	dbUser := "devops"
	dbPass := "devops*123456"
	dbName := "devops"
	dbHost := "10.10.145.80"
	dbPort := "5432"

	tableName := "go_users_user"

	// 创建数据库连接
	db, err := NewDatabase(dbUser, dbPass, dbName, dbHost, dbPort)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}
	defer db.Close()

	fmt.Println("Connected to PostgreSQL!")

	// 执行增删改查操作
	// err := db.InsertData(tableName, "value1", "value2")
	// if err != nil {
	// 	log.Fatal("Insert error:", err)
	// }
	// fmt.Println("Data inserted successfully!")

	rows, err := db.QueryData(tableName)
	if err != nil {
		log.Fatal("Select error:", err)
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var user UserData
		err := rows.Scan(&user.Password,
			&user.LastLogin,
			&user.FirstName,
			&user.LastName,
			&user.IsActive,
			&user.DateJoined,
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.Avatar,
			&user.Wechat,
			&user.Phone,
			&user.PrivateKey,
			&user.PublicKey,
			&user.Comment,
			&user.IsFirstLogin,
			&user.DateExpired,
			&user.CreatedBy,
			&user.MFALevel,
			&user.OTPSecretKey,
			&user.Source,
			&user.DatePasswordLastUpdated,
			&user.NeedUpdatePassword,
			&user.DingtalkID,
			&user.WecomID,
			&user.FeishuID,
			&user.SecretKey,
			&user.IsServiceAccount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User Data:\n%+v\n", user)
	}

	// err = db.UpdateData(tableName, "new_value", "value2")
	// if err != nil {
	// 	log.Fatal("Update error:", err)
	// }
	// fmt.Println("Data updated successfully!")

	// err = db.DeleteData(tableName, "new_value")
	// if err != nil {
	// 	log.Fatal("Delete error:", err)
	// }
	// fmt.Println("Data deleted successfully!")
}
