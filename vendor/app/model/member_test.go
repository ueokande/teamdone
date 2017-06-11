package model

import (
	"app/shared"
	"database/sql"
	"testing"
)

func setFromSlice(s []int64) map[int64]bool {
	set := make(map[int64]bool)
	for _, e := range s {
		set[e] = true
	}
	return set
}

func setupMemberTest() ([]int64, []int64, error) {
	var err error
	orgs := []string{
		"wonderland",
		"glass",
	}
	users := []string{
		"alice",
		"bob",
		"carol",
	}
	members := map[int][]int{
		0: {0, 2},
		1: {1, 2},
	}

	oid := make([]int64, len(orgs), len(orgs))
	for i, name := range orgs {
		oid[i], err = context.OrgCreate(name, shared.RandomKey())
		if err != nil {
			return nil, nil, err
		}
	}

	uid := make([]int64, len(users), len(users))
	for i, name := range users {
		uid[i], err = context.UserCreate(name)
		if err != nil {
			return nil, nil, err
		}
	}

	for o, mem := range members {
		for _, u := range mem {
			err = context.MemberCreate(oid[o], uid[u])
			if err != nil {
				return nil, nil, err
			}
		}
	}
	return oid, uid, nil
}

func TestMemberExists(t *testing.T) {
	oid, uid, err := setupMemberTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	err = context.MemberExists(oid[0], uid[0])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = context.MemberExists(oid[0], uid[2])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = context.MemberExists(oid[1], uid[1])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = context.MemberExists(oid[1], uid[2])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = context.MemberExists(oid[0], uid[1])
	if err != sql.ErrNoRows {
		t.Fatal("Unexpected error:", err)
	}
	err = context.MemberExists(oid[1], uid[0])
	if err != sql.ErrNoRows {
		t.Fatal("Unexpected error:", err)
	}

}

func TestMembersByOrgId(t *testing.T) {
	oid, uid, err := setupMemberTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	member, err := context.MembersByOrgId(oid[0])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if set := setFromSlice(member); !set[uid[0]] || !set[uid[2]] {
		t.Fatal("Unexpected member:", member)
	}

	member, err = context.MembersByOrgId(oid[1])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if set := setFromSlice(member); !set[uid[1]] || !set[uid[2]] {
		t.Fatal("Unexpected member:", member)
	}
}

func TestMembersByUserId(t *testing.T) {
	oid, uid, err := setupMemberTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	member, err := context.MembersByUserId(uid[0])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if set := setFromSlice(member); !set[oid[0]] {
		t.Fatal("Unexpected member:", member)
	}

	member, err = context.MembersByUserId(uid[1])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if set := setFromSlice(member); !set[oid[1]] {
		t.Fatal("Unexpected member:", member)
	}

	member, err = context.MembersByUserId(uid[2])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if set := setFromSlice(member); !set[oid[0]] || !set[oid[1]] {
		t.Fatal("Unexpected member:", member)
	}
}
