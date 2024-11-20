package inventory

type Item struct {
	name     string
	quantity int
}

type Items []Item

func (i *Items) AddItem(newItem *Item) {
	*i = append(*i, *newItem)
}

func (i *Items) RemoveItem(name string) {
	for index, item := range *i {
		if item.name == name {
			*i = append((*i)[:index], (*i)[index+1:]...)
			return
		}
	}
}

func (i *Items) UpdatingQuantity(name string, newQuantity int) {
	for index, item := range *i {
		if item.name == name {
			(*i)[index].quantity = newQuantity
			return
		}
	}
}

func (i *Items) DetailsItem(name string) *Item {
	for _, item := range *i {
		if item.name == name {
			return &item
		}
	}
	return nil
}

func (i *Items) ListItems() []Item {
	return *i
}
