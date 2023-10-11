package impl

import "context"

var (
	_ Interface = (*T)(nil)
)

type T struct{}

func (t *T) Api() {
	panic("not implemented")
}

func (t *T) Api1(ctx context.Context, param Param) Result {
	panic("not implemented")
}

func (t *T) Api2(ctx context.Context, param Param) []Result {
	panic("not implemented")
}

func (t *T) Api3(ctx context.Context, param Param) (r []Result, err error) {
	panic("not implemented")
}

func (t *T) Api4(ctx context.Context, param Param) (r map[*context.Context]Result, err error) {
	panic("not implemented")
}
