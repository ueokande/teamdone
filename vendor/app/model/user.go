package model

import "app/shared/database"

type User struct {
	Id   int64
	Name string
}

func UserCreate(name string) (int64, error) {
	result, err := database.SQL.Exec(
		"INSERT INTO user (name) VALUES (?)",
		name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UserById(id int64) (*User, error) {
	row := database.SQL.QueryRow(
		"SELECT id, name FROM user WHERE id = ? LIMIT 1",
		id)

	var u User
	err := row.Scan(&u.Id, &u.Name)
	if err != nil {
		return nil, err
	}

	return &u, err
}
