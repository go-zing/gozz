// Code generated by gozz:wire. DO NOT EDIT.

//go:build wireinject
// +build wireinject

package inject

import (
	booksallinone "github.com/Just-maple/gozz/examples/books-all-in-one"
	wire "github.com/google/wire"
)

// github.com/Just-maple/gozz/examples/books-all-in-one.Application
func Initialize_booksallinone_Application() (*booksallinone.Application, func(), error) {
	panic(wire.Build(_Set))
}
