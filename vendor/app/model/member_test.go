package model

import (
	"app/shared"
	"database/sql"
	"testing"
)

func setup() ([]int64, []int64, error) {
	var err error
	orgs := []string{
		"wonderland",
		"glass",
	}
	users := []string{
		"alice",
		"bob",
	}
	members := map[int][]int{
		0: {0},
		1: {1},
	}

	oid := make([]int64, len(orgs), len(orgs))
	for i, name := range orgs {
		oid[i], err = OrgCreate(name, shared.RandomKey())
		if err != nil {
			return nil, nil, err
		}
	}
	uid := make([]int64, len(users), len(users))
	for i, name := range users {
		uid[i], err = UserCreate(name)
		if err != nil {
			return nil, nil, err
		}
	}

	for o, mem := range members {
		for _, u := range mem {
			err = MemberCreate(oid[o], uid[u])
			if err != nil {
				return nil, nil, err
			}
		}
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
