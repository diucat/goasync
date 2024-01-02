package scoped

import (
	"runtime/debug"

	"github.com/diucat/goconc/types"
)

func GoFunc[A any](scope Scope, f func() A) func() A {
	return GoFuncWithErr(scope, func() (A, error) {
		return f(), nil
	})
}

func GoFuncWithErr[A any](scope Scope, f func() (A, error)) func() A {
	getter := GoFuncUnchecked(scope, f)

	return func() A {
		val, _, _ := getter()
		return val
	}
}

func GoFuncUnchecked[A any](scope Scope, f func() (A, error)) func() (A, error, bool) {
	idx := scope.total
	scope.total = idx + 1

	result := types.Tup3[A, error, bool]{}

	go func() {
		var (
			val A
			err error
		)

		defer func() {
			if r := recover(); r != nil {
				stackTrace := debug.Stack()
				err = panicError{
					r:          r,
					stackTrace: stackTrace,
				}
			}

			scope.commit(idx, &result.C, types.MkTup[any, error](val, err))
		}()

		val, err = f()
		result.A = val
		result.B = err
	}()

	return func() (A, error, bool) {
		return result.A, result.B, result.C
	}
}

func GoFunc2[A any, B any](scope Scope, f func() (A, B)) func() (A, B) {
	getter := GoFuncWithErr(scope, func() (types.Tup[A, B], error) {
		return types.MkTup(f()), nil
	})

	return func() (A, B) {
		return getter().Unpack()
	}
}

func GoFunc3[A any, B any, C any](scope Scope, f func() (A, B, C)) func() (A, B, C) {
	getter := GoFuncWithErr(scope, func() (types.Tup3[A, B, C], error) {
		return types.MkTup3(f()), nil
	})

	return func() (A, B, C) {
		return getter().Unpack()
	}
}

func GoFunc4[A any, B any, C any, D any](scope Scope, f func() (A, B, C, D)) func() (A, B, C, D) {
	getter := GoFuncWithErr(scope, func() (types.Tup4[A, B, C, D], error) {
		return types.MkTup4(f()), nil
	})

	return func() (A, B, C, D) {
		return getter().Unpack()
	}
}

func GoFunc5[A any, B any, C any, D any, E any](scope Scope, f func() (A, B, C, D, E)) func() (A, B, C, D, E) {
	getter := GoFuncWithErr(scope, func() (types.Tup5[A, B, C, D, E], error) {
		return types.MkTup5(f()), nil
	})

	return func() (A, B, C, D, E) {
		return getter().Unpack()
	}
}

func GoFunc6[A any, B any, C any, D any, E any, F any](scope Scope, f func() (A, B, C, D, E, F)) func() (A, B, C, D, E, F) {
	getter := GoFuncWithErr(scope, func() (types.Tup6[A, B, C, D, E, F], error) {
		return types.MkTup6(f()), nil
	})

	return func() (A, B, C, D, E, F) {
		return getter().Unpack()
	}
}

func GoFunc7[A any, B any, C any, D any, E any, F any, G any](scope Scope, f func() (A, B, C, D, E, F, G)) func() (A, B, C, D, E, F, G) {
	getter := GoFuncWithErr(scope, func() (types.Tup7[A, B, C, D, E, F, G], error) {
		return types.MkTup7(f()), nil
	})

	return func() (A, B, C, D, E, F, G) {
		return getter().Unpack()
	}
}

func GoFunc8[A any, B any, C any, D any, E any, F any, G any, H any](scope Scope, f func() (A, B, C, D, E, F, G, H)) func() (A, B, C, D, E, F, G, H) {
	getter := GoFuncWithErr(scope, func() (types.Tup8[A, B, C, D, E, F, G, H], error) {
		return types.MkTup8(f()), nil
	})

	return func() (A, B, C, D, E, F, G, H) {
		return getter().Unpack()
	}
}

func GoFunc2WithErr[A any, B any](scope Scope, f func() (A, B, error)) func() (A, B) {
	getter := GoFuncWithErr(scope, func() (types.Tup[A, B], error) {
		a, b, err := f()
		return types.MkTup(a, b), err
	})

	return func() (A, B) {
		return getter().Unpack()
	}
}

func GoFunc3WithErr[A any, B any, C any](scope Scope, f func() (A, B, C, error)) func() (A, B, C) {
	getter := GoFuncWithErr(scope, func() (types.Tup3[A, B, C], error) {
		a, b, c, err := f()
		return types.MkTup3(a, b, c), err
	})

	return func() (A, B, C) {
		return getter().Unpack()
	}
}

func GoFunc4WithErr[A any, B any, C any, D any](scope Scope, f func() (A, B, C, D, error)) func() (A, B, C, D) {
	getter := GoFuncWithErr(scope, func() (types.Tup4[A, B, C, D], error) {
		a, b, c, d, err := f()
		return types.MkTup4(a, b, c, d), err
	})

	return func() (A, B, C, D) {
		return getter().Unpack()
	}
}

func GoFunc5WithErr[A any, B any, C any, D any, E any](scope Scope, f func() (A, B, C, D, E, error)) func() (A, B, C, D, E) {
	getter := GoFuncWithErr(scope, func() (types.Tup5[A, B, C, D, E], error) {
		a, b, c, d, e, err := f()
		return types.MkTup5(a, b, c, d, e), err
	})

	return func() (A, B, C, D, E) {
		return getter().Unpack()
	}
}

func GoFunc6WithErr[A any, B any, C any, D any, E any, F any](scope Scope, f func() (A, B, C, D, E, F, error)) func() (A, B, C, D, E, F) {
	getter := GoFuncWithErr(scope, func() (types.Tup6[A, B, C, D, E, F], error) {
		a, b, c, d, e, f, err := f()
		return types.MkTup6(a, b, c, d, e, f), err
	})

	return func() (A, B, C, D, E, F) {
		return getter().Unpack()
	}
}

func GoFunc7WithErr[A any, B any, C any, D any, E any, F any, G any](scope Scope, f func() (A, B, C, D, E, F, G, error)) func() (A, B, C, D, E, F, G) {
	getter := GoFuncWithErr(scope, func() (types.Tup7[A, B, C, D, E, F, G], error) {
		a, b, c, d, e, f, g, err := f()
		return types.MkTup7(a, b, c, d, e, f, g), err
	})

	return func() (A, B, C, D, E, F, G) {
		return getter().Unpack()
	}
}

func GoFunc8WithErr[A any, B any, C any, D any, E any, F any, G any, H any](scope Scope, f func() (A, B, C, D, E, F, G, H, error)) func() (A, B, C, D, E, F, G, H) {
	getter := GoFuncWithErr(scope, func() (types.Tup8[A, B, C, D, E, F, G, H], error) {
		a, b, c, d, e, f, g, h, err := f()
		return types.MkTup8(a, b, c, d, e, f, g, h), err
	})

	return func() (A, B, C, D, E, F, G, H) {
		return getter().Unpack()
	}
}
