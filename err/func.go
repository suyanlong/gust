package err

import (
	"github.com/henrylee2cn/gust"
)

func And[T any](err error, r gust.Result[T]) gust.Result[T] {
	if err != nil {
		return gust.Err[T](err)
	}
	return r
}

func AndThen[T any](err error, op func() gust.Result[T]) gust.Result[T] {
	if err != nil {
		return gust.Err[T](err)
	}
	return op()
}
