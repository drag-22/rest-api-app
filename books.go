package main

var Books = NewBookManager()

type Book struct {
	Id     int
	Author string
	Name   string
}

type BookManager struct {
	list []Book
	pool pool
}

func NewBookManager() *BookManager {
	return &BookManager{pool: makePool(2)}
}

func (b *BookManager) Add(author, name string) {
	b.list = append(b.list, Book{
		Author: author,
		Name:   name,
		Id:     b.pool.get(),
	})
}

func (b *BookManager) Rem(id int) {
	temp := b.list[:0]
	for _, book := range b.list {
		if book.Id != id {
			temp = append(temp, book)
		}
	}
	b.list = temp
	b.pool.rem(id)
}
func (b *BookManager) Reset() {
	b.list = b.list[:0]
	b.pool.reset()
}
