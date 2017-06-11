package model

type Org struct {
	Id   int64
	Name string
	Key  string
}

func (c *Context) OrgCreate(name string, key string) (int64, error) {
	result, err := c.SQL.Exec(
		"INSERT INTO org (name, `key`) VALUES (?, ?)",
		name, key)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (c *Context) OrgById(id int64) (*Org, error) {
	row := c.SQL.QueryRow(
		"SELECT id, name, `key` FROM org WHERE id = ? LIMIT 1",
		id)

	var o Org
	err := row.Scan(&o.Id, &o.Name, &o.Key)
	if err != nil {
		return nil, err
	}

	return &o, err
}

func (c *Context) OrgByKey(key string) (*Org, error) {
	row := c.SQL.QueryRow(
		"SELECT id, name, `key` FROM org WHERE `key` = ? LIMIT 1",
		key)

	var o Org
	err := row.Scan(&o.Id, &o.Name, &o.Key)
	if err != nil {
		return nil, err
	}

	return &o, err
}

func (c *Context) OrgsByUserId(uid int64) ([]*Org, error) {
	rows, err := c.SQL.Query(
		"SELECT id, name, `key` from org INNSER JOIN member ON id = member.org_id WHERE user_id = ?",
		uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orgs := make([]*Org, 0)
	for rows.Next() {
		o := new(Org)
		err := rows.Scan(&o.Id, &o.Name, &o.Key)
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, o)
	}
	return orgs, nil
}
