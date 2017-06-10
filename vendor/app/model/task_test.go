package model

import (
	"app/shared"
	"testing"
	"time"
)

func setupTasks() ([]int64, error) {
	var err error
	orgs := []string{
		"wonderland",
		"glass",
	}
	date := time.Date(2010, time.December, 24, 0, 0, 0, 0, time.UTC)
	tasks1 := []struct {
		summary string
		due     *time.Time
	}{
		{"apple", nil},
		{"banana", &date},
	}

	tasks2 := []string{
		"ximenia",
		"yucca",
	}

	oid := make([]int64, len(orgs), len(orgs))
	for i, name := range orgs {
		oid[i], err = OrgCreate(name, shared.RandomKey())
		if err != nil {
			return nil, err
		}
	}
	for _, t := range tasks1 {
		_, err = TaskCreate(oid[0], t.summary, "", t.due)
		if err != nil {
			return nil, err
		}
	}
	for _, summary := range tasks2 {
		_, err = TaskCreate(oid[1], summary, "", nil)
		if err != nil {
			return nil, err
		}
	}
	return oid, nil
}

func TestTaskCreate(t *testing.T) {
	oid, err := OrgCreate("wonderland", shared.RandomKey())
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

func TestTaskByOrgId(t *testing.T) {
	oid, err := setupTasks()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	tasks, err := TaskByOrgId(oid[0])
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if len(tasks) != 2 ||
		tasks[0].Summary != "apple" || tasks[1].Summary != "banana" ||
		tasks[0].Due != nil || *tasks[1].Due != time.Date(2010, time.December, 24, 0, 0, 0, 0, time.UTC) {
		t.Fatal("Unexpected task len:", tasks)
	}
}
