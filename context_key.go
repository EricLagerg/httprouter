package httprouter

import "golang.org/x/net/context"

type key uint8

const (
	paramsKey key = iota
	reqKey
)

func (p Params) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, paramsKey, p)
}

func FromContext(ctx context.Context) (Params, bool) {
	p, ok := ctx.Value(paramsKey).(Params)
	return p, ok
}

func reqIDContext() context.Context {
	return context.WithValue(context.Background(), reqKey, reqID())
}

func ReqIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(reqKey).(string)
	return id, ok
}
