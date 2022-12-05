package datastructures

type Stack[T any] []T

// IsEmpty  check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// Pop removes and returns top element of stack. Return false if stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return getZero[T](), false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack[T]) Take(n int) (out []T, ok bool) {
	if len(*s) < n {
		return nil, false
	}

	for i := 0; i < n; i++ {
		out = append(out, (*s)[len(*s)-n+i])
	}

	*s = (*s)[:len(*s)-n]

	return out, true
}

func getZero[T any]() T {
	var result T
	return result
}
