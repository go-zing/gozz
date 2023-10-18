// Code generated by gozz:wire. DO NOT EDIT.

package inject

import (
	"context"
	booksallinone "github.com/go-zing/gozz/examples/books-all-in-one"
)

type _aop_interceptor interface {
	Intercept(v interface{}, name string, params, results []interface{}) (func(), bool)
}

// booksallinone.BookService
type (
	_aop_booksallinone_BookService      booksallinone.BookService
	_impl_aop_booksallinone_BookService struct{ _aop_booksallinone_BookService }
)

func (i _impl_aop_booksallinone_BookService) List(p0 context.Context, p1 booksallinone.QueryBook) (r0 booksallinone.ListBook, r1 error) {
	if t, x := i._aop_booksallinone_BookService.(_aop_interceptor); x {
		if up, ok := t.Intercept(i._aop_booksallinone_BookService, "List",
			[]interface{}{&p0, &p1},
			[]interface{}{&r0, &r1},
		); up != nil {
			defer up()
		} else if !ok {
			return
		}
	}
	return i._aop_booksallinone_BookService.List(p0, p1)
}

func (i _impl_aop_booksallinone_BookService) Get(p0 context.Context, p1 booksallinone.QueryBook) (r0 booksallinone.DataBook, r1 error) {
	if t, x := i._aop_booksallinone_BookService.(_aop_interceptor); x {
		if up, ok := t.Intercept(i._aop_booksallinone_BookService, "Get",
			[]interface{}{&p0, &p1},
			[]interface{}{&r0, &r1},
		); up != nil {
			defer up()
		} else if !ok {
			return
		}
	}
	return i._aop_booksallinone_BookService.Get(p0, p1)
}

func (i _impl_aop_booksallinone_BookService) Create(p0 context.Context, p1 booksallinone.FormBook) (r0 booksallinone.DataBook, r1 error) {
	if t, x := i._aop_booksallinone_BookService.(_aop_interceptor); x {
		if up, ok := t.Intercept(i._aop_booksallinone_BookService, "Create",
			[]interface{}{&p0, &p1},
			[]interface{}{&r0, &r1},
		); up != nil {
			defer up()
		} else if !ok {
			return
		}
	}
	return i._aop_booksallinone_BookService.Create(p0, p1)
}

func (i _impl_aop_booksallinone_BookService) Edit(p0 context.Context, p1 booksallinone.FormBook) (r0 booksallinone.DataBook, r1 error) {
	if t, x := i._aop_booksallinone_BookService.(_aop_interceptor); x {
		if up, ok := t.Intercept(i._aop_booksallinone_BookService, "Edit",
			[]interface{}{&p0, &p1},
			[]interface{}{&r0, &r1},
		); up != nil {
			defer up()
		} else if !ok {
			return
		}
	}
	return i._aop_booksallinone_BookService.Edit(p0, p1)
}