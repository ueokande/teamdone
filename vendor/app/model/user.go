package model

type User struct {
	Id   int64
	Name string
}

func (c *Context) UserCreate(name string) (int64, error) {
	result, err := c.SQL.Exec(
		"INSERT INTO user (name) VALUES (?)",
		name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (c *Context) UserById(id int64) (*User, error) {
	row := c.SQL.QueryRow(
		"SELECT id, name FROM user WHERE id = ? LIMIT 1",
		id)

	var u User
	err := row.Scan(&u.Id, &u.Name)
	if err != nil {
		return nil, err
	}

	return &u, err
}
