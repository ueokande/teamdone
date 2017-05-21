package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

var UnknownAdapter = errors.New("unknown adapter")
var SQL *sql.DB

type Database interface {
	DSN() string
}

type MySQL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Name     string `json:"database"`
}

func (info MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		info.Username, info.Password,
		info.Host, info.Port,
		info.Name)
}

func LoadConfig(path string) (Database, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	kvs := make(map[string]interface{})
	err = json.Unmarshal(data, &kvs)
	if err != nil {
		return nil, err
	}

	switch kvs["adapter"] {
	case "mysql":
		db := new(MySQL)
		err := json.Unmarshal(data, db)
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return nil, UnknownAdapter
}

func Connect(db Database) error {
	var err error

	SQL, err = sql.Open("mysql", db.DSN())
	if err != nil {
		return err
	}

	err = SQL.Ping()
	if err != nil {
		return err
	}

	return nil
}
