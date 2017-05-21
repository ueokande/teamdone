package model

import "testing"

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
