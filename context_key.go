package httprouter

import "golang.org/x/net/context"

type key uint8

const paramsKey key = 0

func (p Params) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, paramsKey, p)
}

func FromContext(ctx context.Context) (Params, bool) {
	p, ok := ctx.Value(paramsKey).(Params)
	return p, ok
}
