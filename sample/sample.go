package sample

import (
	"context"

	"github.com/diucat/goconc/scoped"
)

func Sample(ctx context.Context) (int, error) {
	scope := scoped.MakeScope()

	getter := scoped.GoFunc(scope,
		func() int {
			return 1
		})

	getter2 := scoped.GoFunc(scope,
		func() int {
			return 1
		})

	getter3 := scoped.GoFunc(scope,
		func() int {

			scope := scope.Child()

			return 1
		})

	scope.Join()

	if err := scope.Err(); err != nil {
		return 0, err
	}

	return getter() + getter2(), nil
}
