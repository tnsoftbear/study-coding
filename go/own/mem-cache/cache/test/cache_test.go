package test

import (
	"context"
	"encoding/json"
	"testing"
	"time"
	
	"mem-cache/model"
	"mem-cache/cache"
)

func TestAbsence(t *testing.T) {
	const KEY = "absent"
	sut := cache.NewCache(context.TODO(), 5 * time.Millisecond)
	if sut.Has(KEY) {
		t.Errorf("Checking item absence by key '%s', but item exists", KEY)
	}
	v, err := sut.Get(KEY)
	if v != nil {
		t.Errorf("Reading absent item by key '%s', but item is found", KEY)
	}
	if err != cache.ErrNotFound {
		t.Errorf("Reading absent item by key '%s' and expecting 'not found' error", KEY)
	}
}

func TestExistence(t *testing.T) {
	const KEY = "existing"
	sut := cache.NewCache(context.TODO(), 5 * time.Millisecond)
	now := time.Now()
	initialProfile := model.Profile{
		UUID: "id1",
		Name: "name1",
		Orders: []*model.Order{
			{
				UUID:      "id1",
				Value:     "value1",
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				UUID:      "id1",
				Value:     "value2",
				CreatedAt: now.Add(-5 * time.Minute),
				UpdatedAt: now.Add(-3 * time.Minute),
			},
		},
	}
	expectedProfileJSON, _ := json.Marshal(initialProfile)
	err := sut.Set(KEY, initialProfile)
	if err != nil {
		t.Fatalf("Storing item in cache failed with error: %v", err)
	}

	// Изменяем профиль, это не должно влиять на состояния данных в кэше
	initialProfile.Name = "name1_changed"
	initialProfile.Orders[0].Value = "value1_changed"

	if !sut.Has(KEY) {
		t.Errorf("Checking item existence by key '%s', but item is absent", KEY)
	}
	cachedProfile, err := sut.Get(KEY)
	if err != nil {
		t.Errorf("Reading existing item by key '%s', but got error '%s'", KEY, err.Error())
	}

	cachedProfileJSON, _ := json.Marshal(cachedProfile)
	if string(expectedProfileJSON) != string(cachedProfileJSON) {
		t.Errorf("Reading existing item by key '%s', but expected value (%+v) not equal to actual (%+v)", KEY, string(expectedProfileJSON), string(cachedProfileJSON))
	}
}
