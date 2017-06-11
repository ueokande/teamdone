package model

func (c *Context) MemberExists(oid, uid int64) error {
	row := c.SQL.QueryRow(
		"SELECT 1 FROM member WHERE org_id = ? AND user_id = ? limit 1",
		oid, uid)
	var result int64
	return row.Scan(&result)
}

func (c *Context) MemberCreate(oid, uid int64) error {
	_, err := c.SQL.Exec(
		"INSERT INTO member (org_id, user_id) VALUES (?, ?)",
		oid, uid)
	return err
}

func (c *Context) MembersByUserId(uid int64) ([]int64, error) {
	rows, err := c.SQL.Query(
		"SELECT org_id FROM member WHERE user_id = ?",
		uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	oids := make([]int64, 0)
	for rows.Next() {
		var o int64
		err := rows.Scan(&o)
		if err != nil {
			return nil, err
		}
		oids = append(oids, o)
	}
	return oids, nil
}

func (c *Context) MembersByOrgId(oid int64) ([]int64, error) {
	rows, err := c.SQL.Query(
		"SELECT user_id FROM member WHERE org_id = ?",
		oid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	uids := make([]int64, 0)
	for rows.Next() {
		var u int64
		err := rows.Scan(&u)
		if err != nil {
			return nil, err
		}
		uids = append(uids, u)
	}
	return uids, nil
}
