package types

type Tup[A any, B any] struct {
	A A
	B B
}

func MkTup[A any, B any](a A, b B) Tup[A, B] {
	return Tup[A, B]{
		A: a,
		B: b,
	}
}

func (tup Tup[A, B]) Unpack() (A, B) {
	return tup.A, tup.B
}

type Tup3[A any, B any, C any] struct {
	A A
	B B
	C C
}

func MkTup3[A any, B any, C any](a A, b B, c C) Tup3[A, B, C] {
	return Tup3[A, B, C]{
		A: a,
		B: b,
		C: c,
	}
}

func (tup Tup3[A, B, C]) Unpack() (A, B, C) {
	return tup.A, tup.B, tup.C
}

type Tup4[A any, B any, C any, D any] struct {
	A A
	B B
	C C
	D D
}

func MkTup4[A any, B any, C any, D any](a A, b B, c C, d D) Tup4[A, B, C, D] {
	return Tup4[A, B, C, D]{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func (tup Tup4[A, B, C, D]) Unpack() (A, B, C, D) {
	return tup.A, tup.B, tup.C, tup.D
}

type Tup5[A any, B any, C any, D any, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

func MkTup5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Tup5[A, B, C, D, E] {
	return Tup5[A, B, C, D, E]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
	}
}

func (tup Tup5[A, B, C, D, E]) Unpack() (A, B, C, D, E) {
	return tup.A, tup.B, tup.C, tup.D, tup.E
}

type Tup6[A any, B any, C any, D any, E any, F any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

func MkTup6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Tup6[A, B, C, D, E, F] {
	return Tup6[A, B, C, D, E, F]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
	}
}

func (tup Tup6[A, B, C, D, E, F]) Unpack() (A, B, C, D, E, F) {
	return tup.A, tup.B, tup.C, tup.D, tup.E, tup.F
}

type Tup7[A any, B any, C any, D any, E any, F any, G any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

func MkTup7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Tup7[A, B, C, D, E, F, G] {
	return Tup7[A, B, C, D, E, F, G]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
	}
}

func (tup Tup7[A, B, C, D, E, F, G]) Unpack() (A, B, C, D, E, F, G) {
	return tup.A, tup.B, tup.C, tup.D, tup.E, tup.F, tup.G
}

type Tup8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

func MkTup8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tup8[A, B, C, D, E, F, G, H] {
	return Tup8[A, B, C, D, E, F, G, H]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
	}
}

func (tup Tup8[A, B, C, D, E, F, G, H]) Unpack() (A, B, C, D, E, F, G, H) {
	return tup.A, tup.B, tup.C, tup.D, tup.E, tup.F, tup.G, tup.H
}
