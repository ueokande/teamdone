package model

import "app/shared/database"

func MemberExists(oid, uid int64) error {
	row := database.SQL.QueryRow(
		"SELECT 1 FROM member WHERE org_id = ? AND user_id = ? limit 1",
		oid, uid)
	var result int64
	return row.Scan(&result)
}

func MemberCreate(oid, uid int64) error {
	_, err := database.SQL.Exec(
		"INSERT INTO member (org_id, user_id) VALUES (?, ?)",
		oid, uid)
	return err
}
