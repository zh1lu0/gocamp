package dao

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "gocamp",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

/**
* NameById 我觉得错误可以 Wrap 抛给上层，交给业务层来打日志
 */
func NameById(id int64) (string, error) {
	var name string
	row := db.QueryRow("SELECT username FROM users WHERE id = ?", id)
	if err := row.Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return name, fmt.Errorf("NameById %d: no such user", id)
		}
		return name, fmt.Errorf("NameById %d: %v", id, err)
	}
	return name, nil
}
