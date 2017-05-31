package model

import (
	"testing"
)

func TestOrgCreate(t *testing.T) {
	id, err := OrgCreate("wonderland", "abcd1234")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	o, err := OrgById(id)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if o.Name != "wonderland" {
		t.Fatal("Unexpected org name:", o.Name)
	}
	if string(o.Key) != "abcd1234" {
		t.Fatal("Unexpected org key:", o.Key)
	}

}

func TestOrgByKey(t *testing.T) {
	id, err := OrgCreate("wonderland", "secret-key")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	o, err := OrgByKey("secret-key")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if o.Id != id {
		t.Fatal("Unexpected org id:", o.Id)
	}
}

func TestOrgsByUserId(t *testing.T) {
	oid, uid, err := setupMemberTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	orgs, err := OrgsByUserId(uid[0])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if len(orgs) != 1 || orgs[0].Id != oid[0] {
		t.Fatal("Unexpected orgs:", orgs)
	}

	orgs, err = OrgsByUserId(uid[2])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if len(orgs) != 2 || orgs[0].Id != oid[0] || orgs[1].Id != oid[1] {
		t.Fatal("Unexpected orgs:", orgs)
	}

}
