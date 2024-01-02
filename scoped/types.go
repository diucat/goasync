package scoped

import (
	"context"
)

type Scope struct {
	ctx    context.Context
	cancel func()

	chSync chan *bool
	chDone chan struct{}

	total    int
	selected int
	val      interface{}
	err      error
}

type ScopedConc[T any] struct {
	scope Scope
}

type result struct {
}

type panicError struct {
	r          any
	stackTrace []byte
}
