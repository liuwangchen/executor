package executor

import (
	"context"
	"errors"
)

var (
	ErrNonePlan = errors.New("none plan")
)

//Executor interface definition
type Executor interface {
	//Execute before stopping make sure all subroutines stopped
	Execute(context.Context) error
}

//ExecutorFunc definition
type ExecutorFunc func(context.Context) error

//Execute ExecutorFunc implemention of Executor
func (f ExecutorFunc) Execute(ctx context.Context) error {
	return f(ctx)
}

// ExecutorMiddleware is a function that middlewares can implement to be
// able to chain.
type ExecutorMiddleware func(Executor) Executor

// UseMiddleware wraps a Executor in one or more middleware.
func UseMiddleware(exec Executor, middleware ...ExecutorMiddleware) Executor {
	// Apply in reverse order.
	for i := len(middleware) - 1; i >= 0; i-- {
		m := middleware[i]
		exec = m(exec)
	}
	return exec
}
