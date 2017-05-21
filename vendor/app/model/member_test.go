package model

import (
	"app/shared"
	"database/sql"
	"testing"
)

func setup() ([]int64, []int64, error) {
	var err error
	oid := make([]int64, 2, 2)
	uid := make([]int64, 2, 2)

	oid[0], err = OrgCreate("wonderland", shared.RandomKey())
	if err != nil {
		return nil, nil, err
	}
	oid[1], err = OrgCreate("glass", shared.RandomKey())
	if err != nil {
		return nil, nil, err
	}
	uid[0], err = UserCreate("alice")
	if err != nil {
		return nil, nil, err
	}
	uid[1], err = UserCreate("bob")
	if err != nil {
		return nil, nil, err
	}

	err = MemberCreate(oid[0], uid[0])
	if err != nil {
		return nil, nil, err
	}
	err = MemberCreate(oid[1], uid[1])
	if err != nil {
		return nil, nil, err
	}
	return oid, uid, nil
}

func TestMemberExists(t *testing.T) {
	oid, uid, err := setup()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	err = MemberExists(oid[0], uid[0])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = MemberExists(oid[1], uid[1])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = MemberExists(oid[0], uid[1])
	if err != sql.ErrNoRows {
		t.Fatal("Unexpected error:", err)
	}
	err = MemberExists(oid[1], uid[0])
	if err != sql.ErrNoRows {
		t.Fatal("Unexpected error:", err)
	}

}
