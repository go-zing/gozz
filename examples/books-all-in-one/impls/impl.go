package impls

import (
	"context"

	booksallinone "github.com/go-zing/gozz/examples/books-all-in-one"
)

var (
	_ booksallinone.BookService = (*BookServiceImpl)(nil)
)

// +zz:wire:bind=booksallinone.BookService:aop
type BookServiceImpl struct{}

func (impl *BookServiceImpl) List(ctx context.Context, query booksallinone.QueryBook) (ret booksallinone.ListBook, err error) {
	panic("not implemented")
}

func (impl *BookServiceImpl) Get(ctx context.Context, book booksallinone.QueryBook) (data booksallinone.DataBook, err error) {
	panic("not implemented")
}

func (impl *BookServiceImpl) Create(ctx context.Context, book booksallinone.FormBook) (data booksallinone.DataBook, err error) {
	panic("not implemented")
}

func (impl *BookServiceImpl) Edit(ctx context.Context, book booksallinone.FormBook) (data booksallinone.DataBook, err error) {
	panic("not implemented")
}
