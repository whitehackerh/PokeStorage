package entity

type Item struct {
	id   int
	name string
}

func NewItem(id int, name string) Item {
	return Item{
		id:   id,
		name: name,
	}
}

func (i *Item) Id() int {
	return i.id
}

func (i *Item) Name() string {
	return i.name
}
