package pkg

import (
	"context"

	"github.com/Just-maple/gozz/examples/impl"
)

var (
	_ impl.Interface = (*InterfaceImpl)(nil)
)

type InterfaceImpl struct{}

func (impl *InterfaceImpl) Api() {
	panic("not implemented")
}

func (impl *InterfaceImpl) Api1(ctx context.Context, param impl.Param) impl.Result {
	panic("not implemented")
}

func (impl *InterfaceImpl) Api2(ctx context.Context, param impl.Param) []impl.Result {
	panic("not implemented")
}

func (impl *InterfaceImpl) Api3(ctx context.Context, param impl.Param) (r []impl.Result, err error) {
	panic("not implemented")
}

func (impl *InterfaceImpl) Api4(ctx context.Context, param impl.Param) (r map[*context.Context]impl.Result, err error) {
	panic("not implemented")
}
