package balance

import (
	"reflect"
	"testing"

	types "github.com/KodiakAS/balance/types"
)

func TestBalanceItemsFromJson(t *testing.T) {
	got, err := BalanceItemsFromJson("tests/demo.json", 2)
	if err != nil {
		t.Fatalf("BalanceItemsFromJson error: %v", err)
	}
	want := types.BucketList{
		{Items: types.ItemList{{Data: "Hello", Factor: 1}}, Factor: 1},
		{Items: types.ItemList{{Data: "World", Factor: 2}}, Factor: 2},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BalanceItemsFromJson() = %v, want %v", got, want)
	}
}

func TestBalanceItemsFromJson_Error(t *testing.T) {
	if _, err := BalanceItemsFromJson("no_such.json", 1); err == nil {
		t.Errorf("expected error for missing file")
	}
}

func TestBalanceItems(t *testing.T) {
	items := types.ItemList{
		{Data: "1", Factor: 1},
		{Data: "2", Factor: 2},
	}
	got, err := BalanceItems(items, 2)
	if err != nil {
		t.Fatalf("BalanceItems error: %v", err)
	}
	want := types.BucketList{
		{Items: types.ItemList{{Data: "1", Factor: 1}}, Factor: 1},
		{Items: types.ItemList{{Data: "2", Factor: 2}}, Factor: 2},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BalanceItems() = %v, want %v", got, want)
	}
}

func TestBalanceItems_InvalidBucket(t *testing.T) {
	items := types.ItemList{{Data: "1", Factor: 1}}
	if _, err := BalanceItems(items, 0); err == nil {
		t.Errorf("expected error with zero buckets")
	}
}
