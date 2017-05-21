package model

import (
	"testing"
	"time"
)

func TestTaskCreate(t *testing.T) {
	oid, err := OrgCreate("wonderland", "abcdefgh")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	due := time.Date(1865, time.November, 26, 23, 55, 66, 99, time.UTC)
	id, err := TaskCreate(oid, "eat a cake", "written EAT ME", &due)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	task, err := TaskById(id)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if task.OrgId != oid ||
		task.Summary != "eat a cake" ||
		task.Description != "written EAT ME" ||
		task.Done != false {
		t.Fatal("Unexpected task:", task)
	}

}
