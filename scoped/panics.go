package scoped

import "fmt"

func (p panicError) Error() string {
	return fmt.Sprintf("Panic: %+v.\n\n%s", p.r, p.stackTrace)
}
