package inventory

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	items := &Items{}
	item := &Item{name: "Apple", quantity: 10}

	items.AddItem(item)

	if len(*items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(*items))
	}
	if (*items)[0].name != "Apple" {
		t.Errorf("Expected item name 'Apple', got '%s'", (*items)[0].name)
	}
	if (*items)[0].quantity != 10 {
		t.Errorf("Expected item quantity 10, got %d", (*items)[0].quantity)
	}
}

func TestRemoveItem(t *testing.T) {
	items := &Items{
		{name: "Apple", quantity: 10},
		{name: "Banana", quantity: 5},
	}

	items.RemoveItem("Apple")

	if len(*items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(*items))
	}
	if (*items)[0].name != "Banana" {
		t.Errorf("Expected item name 'Banana', got '%s'", (*items)[0].name)
	}
}

func TestUpdatingQuantity(t *testing.T) {
	items := &Items{
		{name: "Apple", quantity: 10},
	}

	items.UpdatingQuantity("Apple", 20)

	if (*items)[0].quantity != 20 {
		t.Errorf("Expected item quantity 20, got %d", (*items)[0].quantity)
	}
}

func TestDetailsItem(t *testing.T) {
	items := &Items{
		{name: "Apple", quantity: 10},
	}

	item := items.DetailsItem("Apple")

	if item == nil {
		t.Error("Expected to find item 'Apple', got nil")
	} else if item.name != "Apple" {
		t.Errorf("Expected item name 'Apple', got '%s'", item.name)
	}
}

func TestListItems(t *testing.T) {
	items := &Items{
		{name: "Apple", quantity: 10},
		{name: "Banana", quantity: 5},
	}

	list := items.ListItems()

	if len(list) != 2 {
		t.Errorf("Expected 2 items, got %d", len(list))
	}
}
