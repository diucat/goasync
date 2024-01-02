package scoped

import "github.com/diucat/goconc/types"

func MakeScope() Scope {

}

func (s *Scope) Child() Scope {

}

func (s *Scope) Join() {
	for x := range s.chSync {
		*x = true
		if s.done() {
			return
		}
	}
}

func (s *Scope) Err() error {

}

func (s *Scope) done() bool {

}

func (s *Scope) commit(idx int, b *bool, result types.Tup[any, error]) {
	select {
	case <-s.chDone:
	case s.chSync <- b:
	}
}
